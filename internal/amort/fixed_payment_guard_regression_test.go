package amort_test

import (
	"testing"

	"task109-loanamort/internal/amort"
	"task109-loanamort/internal/domain"
)

func TestFixedPaymentRejectsNegativeAmortization(t *testing.T) {
	_, err := amort.BuildFixedPayment(amort.ScheduleInput{Outstanding: 100000, PeriodicRateMicro: 200000, Periods: 3, Type: domain.EqualInstallment}, 100)
	if err == nil {
		t.Fatal("expected payment below interest to be rejected")
	}
}
