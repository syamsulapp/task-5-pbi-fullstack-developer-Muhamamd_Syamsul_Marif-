![All Text](https://miro.medium.com/v2/resize:fit:1400/1*Ifpd_HtDiK9u6h68SZgNuA.png)

# introduce

<h5>Go learn adalah source code belajar saya tentang golang programming</h5>

<h5>
untuk menjalankan source code tersebut ada beberapa langkah yang harus di selesaikan:
</h5>

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
