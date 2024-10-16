package models

import (
    "time"
)

type Location struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Address   string  `json:"address"`
}

type NewsCreateReq struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Content     string `json:"content"`
    ImageURL    string `json:"image_url"`
}

type BranchesCreateReq struct {
    Name         string   `json:"name"`
    PhoneNumber  string   `json:"phone_number"`
    OpeningHours string   `json:"opening_hours"`
    Location     Location `json:"location"`
}

type UserCardsCreateReq struct {
    UserID         string `json:"user_id"`
    CardNumber     string `json:"card_number"`
    CardHolderName string `json:"card_holder_name"`
    ExpirationDate string `json:"expiration_date"`
    CVV            string `json:"cvv"`
    Type           string `json:"type"`
}

type UserLocationsCreateReq struct {
    Location Location `json:"location"`
    UserID   string   `json:"user_id"`
}

type CartItemsCreateReq struct {
    UserID    string `json:"user_id"`
    ProductID string `json:"product_id"`
    CartID    string `json:"cart_id"`
    Quantity  int32  `json:"quantity"`
}

type CartItemsUpdateReq struct {
    ID       string `json:"id"`
    Quantity int32  `json:"quantity"`
    UserID   string `json:"user_id"`
}

type CartsCreateReq struct {
    UserID      string `json:"user_id"`
    TotalAmount int64  `json:"total_amount"`
}

type CartsUpdateReq struct {
    ID          string `json:"id"`
    TotalAmount int64  `json:"total_amount"`
    UserID      string `json:"user_id"`
}

type CategoriesCreateReq struct {
    Name string `json:"name"`
}

type CategoriesRes struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type OrderItemsCreateReq struct {
    OrderID   string `json:"order_id"`
    ProductID string `json:"product_id"`
    Quantity  int32  `json:"quantity"`
}

type OrdersCreateReq struct {
    UserID       string   `json:"user_id"`
    CourierID    string   `json:"courier_id"`
    PromocodeID  string   `json:"promocode_id"`
    Status       string   `json:"status"`
    DeliveryType string   `json:"delivery_type"`
    PaymentType  string   `json:"payment_type"`
    TotalAmount  int32    `json:"total_amount"`
    Location     Location `json:"location"`
}

type OrdersUpdateReq struct {
    ID        string `json:"id"`
    CourierID string `json:"courier_id"`
    Status    string `json:"status"`
}

type ProductsCreateReq struct {
    CategoryID  string `json:"category_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int32  `json:"price"`
    ImageURL    string `json:"image_url"`
}

type ProductsUpdate struct {
    ID          string `json:"id"`
    CategoryID  string `json:"category_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int32  `json:"price"`
    ImageURL    string `json:"image_url"`
}

type PromocodeUsagesCreateReq struct {
    UserID      string `json:"user_id"`
    PromocodeID string `json:"promocode_id"`
    UsageCount  int32  `json:"usage_count"`
}

type PromocodesCreateReq struct {
    Code          string    `json:"code"`
    DiscountType  string    `json:"discount_type"`
    DiscountValue int32     `json:"discount_value"`
    UsageLimit    int32     `json:"usage_limit"`
    StartDate     time.Time `json:"start_date"`      
    EndDate       time.Time `json:"end_date"`        
    MinOrderValue int32     `json:"min_order_value"`
    IsPublic      bool      `json:"is_public"`
    UserID        string    `json:"user_id"`
}

type PromocodesUpdate struct {
    ID            string    `json:"id"`
    Code          string    `json:"code"`
    DiscountType  string    `json:"discount_type"`
    DiscountValue int32     `json:"discount_value"`
    UsageLimit    int32     `json:"usage_limit"`
    StartDate     time.Time `json:"start_date"`      
    EndDate       time.Time `json:"end_date"`        
    MinOrderValue int32     `json:"min_order_value"`
    IsPublic      bool      `json:"is_public"`
    UserID        string    `json:"user_id"`
}
