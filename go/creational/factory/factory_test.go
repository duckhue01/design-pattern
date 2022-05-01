package factory

import (
	"testing"
)

// 1. to have a common method for every payment method called by Pay
// 2. to be able to delegate the creation of payments methods to the Factory
// 3. to be able to add more payment methods to the library by just adding it to te factory method

func TestGetPaymentMethod(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name    string
		args    args
		want    PaymentMethod
		wantErr bool
	}{
		{
			name: "test case #1: want to get Cast  payment method",
			args: args{
				m: Cash,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test case #2: want to get Debit Card payment method",
			args: args{
				m: DebitCard,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test case #2: want to get Debit Card payment method",
			args: args{
				m: 3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetPaymentMethod(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaymentMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}

func TestCashPM_Pay(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		cp   *CashPM
		args args
		want string
	}{
		{
			name: "test case #1: want to receive correct payment method",
			cp:   &CashPM{},
			args: args{amount: 10},
			want: "cash payment method",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &CashPM{}
			if got := cp.Pay(tt.args.amount); got != tt.want {
				t.Errorf("CashPM.Pay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebitCardPM_Pay(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		dcp  *DebitCardPM
		args args
		want string
	}{
		{
			name: "",
			dcp:  &DebitCardPM{},
			args: args{amount: 10},
			want: "debit card payment method",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dcp := &DebitCardPM{}
			if got := dcp.Pay(tt.args.amount); got != tt.want {
				t.Errorf("DebitCardPM.Pay() = %v, want %v", got, tt.want)
			}
		})
	}
}
