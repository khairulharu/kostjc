package wcAuth

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"kostjc/model/mAuth/rqAuth"
)

func (u *UsersMutator) SetEncryptedPassword(password string, now int64) {
	u.SetPassword(S.EncryptPassword(password))
	u.SetPasswordSetAt(now)
}

func (s *SessionsMutator) ForceLogoutAll(userId uint64, now int64) (removed []*rqAuth.Sessions, errStr string) {
	activeSession := s.AllActiveSession(userId, now)
	query := `-- ` + L.CallerInfo().String() + `
UPDATE ` + s.SqlTableName() + `
SET ` + s.SqlExpiredAt() + ` = ` + I.ToS(now) + `
WHERE ` + s.SqlUserId() + ` = ` + I.UToS(userId) + ` 
	AND ` + s.SqlExpiredAt() + ` < ` + I.ToS(now)
	out := s.Adapter.ExecSql(query)
	return activeSession, X.ToS(out[`error`])
}

func (u *UsersMutator) SetGenUniqueUsernameByEmail(email string, suffix int64) string {
	newUsername := S.Right(S.ValidateIdent(S.LeftOf(email, `@`))+`_`+I.ToS(suffix%10000), 32)
	u.SetUserName(newUsername)
	return newUsername
}
