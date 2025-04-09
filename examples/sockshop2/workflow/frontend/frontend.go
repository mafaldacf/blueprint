// Package Frontend implements the SockShop Frontend service, typically deployed via HTTP
package frontend

import (
	"context"

	"github.com/google/uuid"

	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/cart"
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/catalogue"
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/order"
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/user"
)

type FrontendService interface {
	// List items in cart for current logged in user, or for the current session if not logged in.
	// SessionID can be the empty string for a non-logged in user / new session
	//GetCart(ctx context.Context, sessionID string) ([]cart.Item, error)

	// Deletes the entire cart for a user/session
	DeleteCart(ctx context.Context, sessionID string) error

	// Removes an item from the user/session's cart
	RemoveItem(ctx context.Context, sessionID string, itemID string) error

	// Update item quantity in the user/session's cart
	// If there is no user or session, then a session is created and the sessionID is returned.
	UpdateItem(ctx context.Context, sessionID string, itemID string, quantity int) (newSessionID string, err error)

	// List socks that match any of the tags specified.  Sort the results by the specified database column.
	// order can be "" in which case the default order is used.
	// pageNum is 1-indexed
	// then return a subset of the results.
	ListItems(ctx context.Context, tags []string, order string, pageNum, pageSize int) ([]catalogue.Sock, error)

	// Gets details about a [Sock]
	GetSock(ctx context.Context, itemID string) (catalogue.Sock, error)

	// Lists all tags
	ListTags(ctx context.Context) ([]string, error)

	// Place an order for the specified items
	NewOrder(ctx context.Context, userID, addressID, cardID, cartID string) (order.Order, error)

	// Get all orders for a customer, sorted by date
	GetOrders(ctx context.Context, userID string) ([]order.Order, error)

	// Get an order by ID
	GetOrder(ctx context.Context, orderID string) (order.Order, error)

	// Log in to an existing user account.  Returns an error if the password
	// doesn't match the registered password
	// Returns the new session ID, which will be the user ID of the logged in user.
	Login(ctx context.Context, sessionID, username, password string) (newSessionID string, u user.User, err error)

	// Register a new user account
	// Returns the new session ID, which will be the user ID of the registered user.
	Register(ctx context.Context, sessionID, username, password, email, first, last string) (newSessionID string, err error)

	// Look up a user by customer ID
	GetUser(ctx context.Context, userID string) (user.User, error)

	// Look up an address by address ID
	GetAddress(ctx context.Context, addressID string) (user.Address, error)

	// Adds a new address for a customer
	PostAddress(ctx context.Context, userID string, address user.Address) (string, error)

	// Look up a card by card id.
	GetCard(ctx context.Context, cardID string) (user.Card, error)

	// Adds a new card for a customer
	PostCard(ctx context.Context, userID string, card user.Card) (string, error)

	// Adds an item to the user/session's cart.
	// If there is no user or session, then a session is created and the sessionID is returned.
	AddItem(ctx context.Context, sessionID string, itemID string) (newSessionID string, err error)

	// Loads the catalogue in the catalogue service
	LoadCatalogue(ctx context.Context) (string, error)

	DeleteSock(ctx context.Context, id string) error
}

type FrontendImpl struct {
	user      user.UserService
	catalogue catalogue.CatalogueService
	cart      cart.CartService
	order     order.OrderService
}

// Instantiates the Frontend service, which makes calls to the user, catalogue, cart, and order services
func NewFrontend(ctx context.Context, user user.UserService, catalogue catalogue.CatalogueService, cart cart.CartService, order order.OrderService) (FrontendService, error) {
	f := &FrontendImpl{
		user:      user,
		catalogue: catalogue,
		cart:      cart,
		order:     order,
	}
	return f, nil
}

// AddItem implements Frontend.
func (f *FrontendImpl) AddItem(ctx context.Context, sessionID string, itemID string) (string, error) {
	if sessionID == "" {
		sessionID = uuid.NewString()
	}

	sock, err := f.catalogue.Get(ctx, itemID)
	if err != nil {
		return sessionID, err
	}

	//FIXME: should pass sock.ID instead of itemID but analyzer is not detecting FOREIGN KEY!!
	_, err = f.cart.AddItem(ctx, sessionID, cart.Item{ID: itemID, Quantity: 1, UnitPrice: sock.Price})
	return sessionID, err
}

// RemoteItem implements FrontendImpl.
func (f *FrontendImpl) RemoveItem(ctx context.Context, sessionID string, itemID string) error {
	if sessionID == "" {
		return nil
	}

	return f.cart.RemoveItem(ctx, sessionID, itemID)
}

// GetCart implements FrontendImpl.
func (f *FrontendImpl) GetCart(ctx context.Context, sessionID string) ([]cart.Item, error) {
	if sessionID == "" {
		return nil, nil
	}

	return f.cart.GetCart(ctx, sessionID)
}

// DeleteCart implements FrontendImpl.
func (f *FrontendImpl) DeleteCart(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return nil
	}

	return f.cart.DeleteCart(ctx, sessionID)
}

// GetUser implements FrontendImpl.
func (f *FrontendImpl) GetUser(ctx context.Context, userID string) (user.User, error) {
	f.user.GetUsers(ctx, userID)
	return user.User{}, nil
}

// GetAddresses implements FrontendImpl.
func (f *FrontendImpl) GetAddress(ctx context.Context, addressID string) (user.Address, error) {
	f.user.GetAddresses(ctx, addressID)
	return user.Address{}, nil
}

// GetCards implements FrontendImpl.
func (f *FrontendImpl) GetCard(ctx context.Context, cardID string) (user.Card, error) {
	f.user.GetCards(ctx, cardID)
	return user.Card{}, nil
}

// GetOrder implements FrontendImpl.
func (f *FrontendImpl) GetOrder(ctx context.Context, orderID string) (order.Order, error) {
	return f.order.GetOrder(ctx, orderID)
}

// GetOrders implements FrontendImpl.
func (f *FrontendImpl) GetOrders(ctx context.Context, userID string) ([]order.Order, error) {
	return f.order.GetOrders(ctx, userID)
}

// GetSock implements FrontendImpl.
func (f *FrontendImpl) GetSock(ctx context.Context, itemID string) (catalogue.Sock, error) {
	return f.catalogue.Get(ctx, itemID)
}

// ListItems implements FrontendImpl.
func (f *FrontendImpl) ListItems(ctx context.Context, tags []string, order string, pageNum int, pageSize int) ([]catalogue.Sock, error) {
	return f.catalogue.List(ctx, tags, order, pageNum, pageSize)
}

// ListTags implements FrontendImpl.
func (f *FrontendImpl) ListTags(ctx context.Context) ([]string, error) {
	return f.catalogue.Tags(ctx)
}

// Login implements FrontendImpl.  Merges the session into the user, and returns the user ID
func (f *FrontendImpl) Login(ctx context.Context, sessionID string, username string, password string) (string, user.User, error) {
	u, err := f.user.Login(ctx, username, password)
	if err != nil {
		return sessionID, user.User{}, err
	}

	if sessionID != "" {
		if err := f.cart.MergeCarts(ctx, u.UserID, sessionID); err != nil {
			return u.UserID, u, err
		}
	}

	return u.UserID, u, nil
}

// NewOrder implements FrontendImpl.
func (f *FrontendImpl) NewOrder(ctx context.Context, userID string, addressID string, cardID string, cartID string) (order.Order, error) {
	return f.order.NewOrder(ctx, userID, addressID, cardID, cartID)
}

// PostAddress implements FrontendImpl.
func (f *FrontendImpl) PostAddress(ctx context.Context, userID string, address user.Address) (string, error) {
	return f.user.PostAddress(ctx, userID, address)
}

// PostCard implements FrontendImpl.
func (f *FrontendImpl) PostCard(ctx context.Context, userID string, card user.Card) (string, error) {
	return f.user.PostCard(ctx, userID, card)
}

// Register implements FrontendImpl.
func (f *FrontendImpl) Register(ctx context.Context, sessionID string, username string, password string, email string, first string, last string) (string, error) {
	userID, err := f.user.Register(ctx, username, password, email, first, last)
	if err != nil {
		return sessionID, err
	}

	if sessionID != "" {
		return userID, f.cart.MergeCarts(ctx, userID, sessionID)
	} else {
		return userID, nil
	}
}

// UpdateItem implements FrontendImpl.
func (f *FrontendImpl) UpdateItem(ctx context.Context, sessionID string, itemID string, quantity int) (string, error) {
	item, err := f.catalogue.Get(ctx, itemID)
	if err != nil {
		return sessionID, err
	}

	return sessionID, f.cart.UpdateItem(ctx, sessionID, cart.Item{ID: item.ID, Quantity: quantity, UnitPrice: item.Price})
}

func (f *FrontendImpl) LoadCatalogue(ctx context.Context) (string, error) {
	sock := catalogue.Sock{}
	f.catalogue.AddSock(ctx, sock)
	return "Load catalogue successful", nil
}

func (f *FrontendImpl) DeleteSock(ctx context.Context, id string) error {
	return f.catalogue.DeleteSock(ctx, id)
}
