package farm

import (
	"github.com/SimonRichardson/coherence/pkg/selectors"
	"github.com/pkg/errors"
)

type nop struct{}

// NewNop creates a new nop farm
func NewNop() Farm {
	return nop{}
}

func (nop) Insert(key selectors.Key,
	members []selectors.FieldValueScore,
	quorum selectors.Quorum,
) (selectors.ChangeSet, error) {
	return selectors.ChangeSet{
		Success: make([]selectors.Field, 0),
		Failure: extractFields(members),
	}, nil
}
func (nop) Delete(key selectors.Key,
	members []selectors.FieldValueScore,
	quorum selectors.Quorum,
) (selectors.ChangeSet, error) {
	return selectors.ChangeSet{
		Success: make([]selectors.Field, 0),
		Failure: extractFields(members),
	}, nil
}
func (nop) Select(selectors.Key, selectors.Field, selectors.Quorum) (selectors.FieldValueScore, error) {
	return selectors.FieldValueScore{}, selectors.NewNotFoundError(errors.New("not found"))
}
func (nop) Keys() ([]selectors.Key, error)                   { return nil, nil }
func (nop) Size(selectors.Key) (int64, error)                { return -1, nil }
func (nop) Members(selectors.Key) ([]selectors.Field, error) { return nil, nil }
func (nop) Score(selectors.Key, selectors.Field) (selectors.Presence, error) {
	return selectors.Presence{}, nil
}
func (nop) Repair([]selectors.KeyFieldValue) error { return nil }

func extractFields(members []selectors.FieldValueScore) []selectors.Field {
	res := make([]selectors.Field, len(members))
	for k, v := range members {
		res[k] = v.Field
	}
	return res
}
