# introduce

<h4>Go learn adalah source code belajar saya tentang golang programming</h4>

<h4>
untuk menjalankan source code tersebut ada beberapa langkah yang harus di selesaikan:
</h4>

<br>

# 1. install go language

```Bash
https://go.dev/doc/install
```

# 2. install nodemon for load automaticly server go

```Bash
npm install -g nodemon
```

# 3. Start server go with nodemon

```Bash
nodemon --exec go run namafile.go --signal SIGTERM
```

# 4. Test Endpoint API

```Bash
#test via postman/insomnia

#baseUrl
localhost:8080 -> sesuaikan dengan base url kalian

#endpoint test
{{base_url}}/universitas ->GET
#register
{{base_url}}/fakultas ->GET
#logout
{{base_url}}/mahasiswa ->GET

```
