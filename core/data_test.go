package core_test

import (
	"context"

	"github.com/google/uuid"
)

const email1 = "user@email.com"
const jwt1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI"
const password1 = "$ome.pa$$word123"
const password2 = "anoth#r.passw0rd"

var context1 = context.Background()
var hash1 = []byte("$2a$10$gYXXJulMpoUalXFgmOpKbO6v.nigV2lWf/Z3EwgykLdGzekwGfAbW")
var uuid1, _ = uuid.Parse("625af883-21ee-40d3-bc40-a753cece2f60")
var uuid2, _ = uuid.Parse("bd2c6d0a-1a6f-4e13-96f7-1ae3d9d48187")
