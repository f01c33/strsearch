package strsearch

import (
	"testing"
)

func TestFind(t *testing.T) {
	text := "The ordinance or emergency executive order has to be cleared by both Houses of Parliament within six weeks from today."
	pattern := "Parliament"

	position, end := FindString(text, pattern)
	if position != 79 && end == position+len(pattern) {
		t.Errorf("Expected %d, got %d", 79, position)
	}

	text = "British actor Eddie Redmayne has won the best actor Oscar for The Theory of Everything, while Julianne Moore picked up best actress for Still Alice."
	pattern = "home"

	position, _ = FindString(text, pattern)
	if position != -1 {
		t.Errorf("Expected %d, got %d", -1, position)
	}
}

func TestFindAll(t *testing.T) {
	text := `Be thou blest, Bertram, and succeed thy father
            In manners, as in shape! thy blood and virtue
            Contend for empire in thee, and thy goodness
            Share with thy birthright! Love all, trust a few,
            Do wrong to none: be able for thine enemy
            Rather in power than use, and keep thy friend
            Under thy own life's key: be cheque'd for silence,
            But never tax'd for speech. What heaven more will,
            That thee may furnish and my prayers pluck down,
            Fall on thy head! Farewell, my lord;
            'Tis an unseason'd courtier; good my lord,
            Advise him.`
	pattern := "and"
	idxs := FinadAllString(text, pattern)
	if len(idxs) != 5 {
		t.Errorf("Expected 5 got %d", len(idxs))
	}
	for _, idx := range idxs {
		if text[idx[0]:idx[1]] != pattern {
			t.Errorf("Expected %s got %s", pattern, text[idx[0]:idx[1]])
		}
	}
}

func TestFindBytes(t *testing.T) {
	text := []byte("The ordinance or emergency executive order has to be cleared by both Houses of Parliament within six weeks from today.")
	pattern := []byte("Parliament")

	position, end := Find(text, pattern)
	if position != 79 && end == position+len(pattern) {
		t.Errorf("Expected %d, got %d", 79, position)
	}

	text = []byte("British actor Eddie Redmayne has won the best actor Oscar for The Theory of Everything, while Julianne Moore picked up best actress for Still Alice.")
	pattern = []byte("home")

	position, _ = Find(text, pattern)
	if position != -1 {
		t.Errorf("Expected %d, got %d", -1, position)
	}
}
