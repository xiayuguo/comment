db.createUser(
    {
	"user": "comment", 
	"pwd": "comment", 
	"roles":[{"role": "root", "db": "admin"}], 
	"mechanisms": ["SCRAM-SHA-1"]
    }
)
