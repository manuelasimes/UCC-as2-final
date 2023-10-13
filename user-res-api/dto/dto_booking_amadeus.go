package dto 

type AmadeusRequest struct {
    Data struct {
        OfferID string `json:"offerId"`
        Guests  []struct {
            Name struct {
                Title     string `json:"title"`
                FirstName string `json:"firstName"`
                LastName  string `json:"lastName"`
            } `json:"name"`
            Contact struct {
                Phone string `json:"phone"`
                Email string `json:"email"`
            } `json:"contact"`
        } `json:"guests"`
        Payments []struct {
            Method string `json:"method"`
            Card   struct {
                VendorCode  string `json:"vendorCode"`
                CardNumber  string `json:"cardNumber"`
                ExpiryDate  string `json:"expiryDate"`
            } `json:"card"`
        } `json:"payments"`
    } `json:"data"`
}

