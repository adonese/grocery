package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCart_save(t *testing.T) {
	type fields struct {
		UserID    int
		ProductID int
		Quantity  int
		Token     string
	}
	f := fields{
		UserID:    13,
		ProductID: 32,
		Quantity:  1,
		Token:     "3232232",
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		count   int
	}{
		// TODO: Add test cases.
		{"successful unveiling", f, true, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				UserID:    tt.fields.UserID,
				ProductID: tt.fields.ProductID,
				Quantity:  tt.fields.Quantity,
				Token:     tt.fields.Token,
			}
			if count, err := c.save(); err != nil && count > 0 {
				t.Errorf("Cart.save() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("Cart.save() retrieved_count = %v, provided_count is %v", count, tt.count)

			}
		})
	}
}

func TestCartItems_populate(t *testing.T) {
	type fields struct {
		ID        int
		UserID    int
		CartID    int
		ProductID int
		Quantity  int
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"adding cart to cartitems", fields{UserID: 32, CartID: 32, ProductID: 32, Quantity: 32}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CartItems{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CartID:    tt.fields.CartID,
				ProductID: tt.fields.ProductID,
				Quantity:  tt.fields.Quantity,
			}
			if err := c.populate(); err != nil {
				t.Logf("The error is: %v", err)
				t.Errorf("CartItems.populate() error = %v", err)
			}
		})
	}
}

func TestCartItems_all(t *testing.T) {
	type fields struct {
		ID        int
		UserID    int
		CartID    int
		ProductID int
		Quantity  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []CartItems
		wantErr bool
	}{
		{"get all items in cartitems", fields{}, []CartItems{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CartItems{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CartID:    tt.fields.CartID,
				ProductID: tt.fields.ProductID,
				Quantity:  tt.fields.Quantity,
			}
			_, err := c.all()
			if err != nil {
				t.Errorf("CartItems.all() error =  %v", err)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CartItems.all() = %v, want %v", got, tt.want)
			// }
		})
	}
}
