package printful

type Response struct {
	Code   int `json:"code"`
	Paging struct {
		Total  int `json:"total"`
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}
}

type ResponseProduct struct {
	*Response
	Result []Product `json:"result"`
}

type ResponseVariants struct {
	*Response
	Result Variants `json:"result"`
}

type Product struct {
	ID         int    `json:"id"`
	ExternalID string `json:"external_id"`
	Name       string `json:"name"`
	Variants   int    `json:"variants"`
	Synced     int    `json:"synced"`
}

type Variants struct {
	SyncProduct  Product `json:"sync_product"`
	SyncVariants []struct {
		ID            int    `json:"id"`
		ExternalID    string `json:"external_id"`
		SyncProductID int    `json:"sync_product_id"`
		Name          string `json:"name"`
		Synced        bool   `json:"synced"`
		VariantID     int    `json:"variant_id"`
		RetailPrice   string `json:"retail_price"`
		Currency      string `json:"currency"`
		Product       struct {
			VariantID int    `json:"variant_id"`
			ProductID int    `json:"product_id"`
			Image     string `json:"image"`
			Name      string `json:"name"`
		} `json:"product"`
		Files []struct {
			ID           int    `json:"id"`
			Type         string `json:"type"`
			Hash         string `json:"hash"`
			URL          string `json:"url"`
			Filename     string `json:"filename"`
			MimeType     string `json:"mime_type"`
			Size         int    `json:"size"`
			Width        int    `json:"width"`
			Height       int    `json:"height"`
			DPI          int    `json:"dpi"`
			Status       string `json:"status"`
			Created      int    `json:"created"`
			ThumbnailURL string `json:"thumbnail_url"`
			PreviewURL   string `json:"preview_url"`
			Visible      bool   `json:"visible"`
		} `json:"files"`
		Options []struct {
			ID    int         `json:"id"`
			Value interface{} `json:"value"`
		} `json:"options"`
	} `json:"sync_variants"`
}
