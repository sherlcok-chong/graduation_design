package query

import "context"

const KeyEmail = "KeyEmail"

// AddEmails adds emails to the set.
func (q *Queries) AddEmails(c context.Context, emails ...string) error {
	if len(emails) == 0 {
		return nil
	}
	data := make([]interface{}, len(emails))
	for i, email := range emails {
		data[i] = email
	}
	return q.rdb.SAdd(c, KeyEmail, data...).Err()
}
