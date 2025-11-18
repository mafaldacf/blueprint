package specs

import (
	"github.com/blueprint-uservices/blueprint/blueprint/pkg/wiring"
	"github.com/blueprint-uservices/blueprint/examples/dsb_media_nosql/workflow/mediamicroservices_nosql"
	"github.com/blueprint-uservices/blueprint/plugins/cmdbuilder"
	"github.com/blueprint-uservices/blueprint/plugins/goproc"
	"github.com/blueprint-uservices/blueprint/plugins/http"
	"github.com/blueprint-uservices/blueprint/plugins/linuxcontainer"
	"github.com/blueprint-uservices/blueprint/plugins/memcached"
	"github.com/blueprint-uservices/blueprint/plugins/mongodb"
	"github.com/blueprint-uservices/blueprint/plugins/redis"
	"github.com/blueprint-uservices/blueprint/plugins/thrift"
	"github.com/blueprint-uservices/blueprint/plugins/workflow"
)

var Docker = cmdbuilder.SpecOption{
	Name:        "docker",
	Description: "Deploys each service in a separate container with thrift, and uses mongodb as NoSQL database backends.",
	Build:       makeDockerSpec,
}

func makeDockerSpec(spec wiring.WiringSpec) ([]string, error) {
	var containers []string

	compose_review_cache := memcached.Container(spec, "compose_review_cache")
	user_review_db := mongodb.Container(spec, "user_review_db")
	user_review_cache := memcached.Container(spec, "user_review_cache")
	movieid_db := mongodb.Container(spec, "movie_id_db")
	movieid_cache := memcached.Container(spec, "movie_id_cache")
	review_storage_db := mongodb.Container(spec, "review_storage_db")
	review_storage_cache := memcached.Container(spec, "review_storage_cache")
	movie_review_db := mongodb.Container(spec, "movie_review_db")
	movie_review_cache := memcached.Container(spec, "movie_review_cache")
	user_db := mongodb.Container(spec, "user_db")
	user_cache := memcached.Container(spec, "user_cache")
	movieinfo_db := mongodb.Container(spec, "movie_info_db")
	castinfo_db := mongodb.Container(spec, "cast_info_db")
	plot_db := mongodb.Container(spec, "plot_db")
	rating_cache := redis.Container(spec, "rating_cache")

	review_storage_service := workflow.Service[mediamicroservices_nosql.ReviewStorageService](spec, "review_storage_service", review_storage_db, review_storage_cache)
	review_storage_ctr := applyDockerDefaults(spec, review_storage_service, "review_storage_ctr", "review_storage_container")
	containers = append(containers, review_storage_ctr)

	user_review_service := workflow.Service[mediamicroservices_nosql.UserReviewService](spec, "user_review_service", user_review_db, user_review_cache, review_storage_service)
	user_review_ctr := applyDockerDefaults(spec, user_review_service, "user_review_ctr", "user_review_container")
	containers = append(containers, user_review_ctr)

	movie_review_service := workflow.Service[mediamicroservices_nosql.MovieReviewService](spec, "movie_review_service", movie_review_db, movie_review_cache, review_storage_service)
	movie_review_ctr := applyDockerDefaults(spec, movie_review_service, "movie_review_ctr", "movie_review_container")
	containers = append(containers, movie_review_ctr)

	compose_review_service := workflow.Service[mediamicroservices_nosql.ComposeReviewService](spec, "compose_review_service", compose_review_cache, review_storage_service, user_review_service, movie_review_service)
	compose_review_ctr := applyDockerDefaults(spec, compose_review_service, "compose_review_proc", "compose_review_container")
	containers = append(containers, compose_review_ctr)

	rating_service := workflow.Service[mediamicroservices_nosql.RatingService](spec, "rating_service", rating_cache, compose_review_service)
	rating_ctr := applyDockerDefaults(spec, rating_service, "rating_ctr", "rating_container")
	containers = append(containers, rating_ctr)

	text_service := workflow.Service[mediamicroservices_nosql.TextService](spec, "text_service", compose_review_service)
	text_ctr := applyDockerDefaults(spec, text_service, "text_ctr", "text_container")
	containers = append(containers, text_ctr)
	
	unique_id_service := workflow.Service[mediamicroservices_nosql.UniqueIdService](spec, "unique_id_service", compose_review_service)
	unique_id_ctr := applyDockerDefaults(spec, unique_id_service, "unique_id_ctr", "unique_id_container")
	containers = append(containers, unique_id_ctr)

	user_service := workflow.Service[mediamicroservices_nosql.UserService](spec, "user_service", user_db, user_cache, compose_review_service)
	user_ctr := applyDockerDefaults(spec, user_service, "user_ctr", "user_container")
	containers = append(containers, user_ctr)

	movieid_service := workflow.Service[mediamicroservices_nosql.MovieIdService](spec, "movieid_service", movieid_db, movieid_cache, compose_review_service, rating_service)
	movieid_ctr := applyDockerDefaults(spec, movieid_service, "movieid_proc", "movieid_container")
	containers = append(containers, movieid_ctr)

	movieinfo_service := workflow.Service[mediamicroservices_nosql.MovieInfoService](spec, "movieinfo_service", movieinfo_db)
	movieinfo_ctr := applyDockerDefaults(spec, movieinfo_service, "movieinfo_proc", "movieinfo_container")
	containers = append(containers, movieinfo_ctr)

	castinfo_service := workflow.Service[mediamicroservices_nosql.CastInfoService](spec, "castinfo_service", castinfo_db)
	castinfo_ctr := applyDockerDefaults(spec, castinfo_service, "castinfo_proc", "castinfo_container")
	containers = append(containers, castinfo_ctr)

	plot_service := workflow.Service[mediamicroservices_nosql.PlotService](spec, "plot_service", plot_db)
	plot_ctr := applyDockerDefaults(spec, plot_service, "plot_ctr", "plot_container")
	containers = append(containers, plot_ctr)

	page_service := workflow.Service[mediamicroservices_nosql.PageService](spec, "page_service", movieinfo_service, movie_review_service, castinfo_service, plot_service)
	page_ctr := applyDockerDefaults(spec, page_service, "page_proc", "page_container")
	containers = append(containers, page_ctr)

	api_service := workflow.Service[mediamicroservices_nosql.APIService](spec, "api_service", user_service, text_service, movieid_service, unique_id_service, movieinfo_service, castinfo_service, plot_service, page_service)
	api_ctr := applyHTTPDefaults(spec, api_service, "api_proc", "api_container")
	containers = append(containers, api_ctr)

	return containers, nil
}

func applyDockerDefaults(spec wiring.WiringSpec, serviceName, procName, ctrName string) string {
	thrift.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}

func applyHTTPDefaults(spec wiring.WiringSpec, serviceName, procName, ctrName string) string {
	http.Deploy(spec, serviceName)
	goproc.CreateProcess(spec, procName, serviceName)
	return linuxcontainer.CreateContainer(spec, ctrName, procName)
}
