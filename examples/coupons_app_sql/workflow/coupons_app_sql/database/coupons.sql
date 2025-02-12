CREATE TABLE IF NOT EXISTS coupons (
	coupon_id   INT NOT NULL, 
	category    VARCHAR(20),
	PRIMARY KEY(coupon_id)
);

CREATE TABLE IF NOT EXISTS claimed_coupons (
    coupon_id   INT NOT NULL, 
    user_id     INT NOT NULL,
    PRIMARY KEY(coupon_id, user_id)
);
