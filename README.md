# gRPC Student Service (Golang)

โปรเจกต์นี้เป็นตัวอย่างการพัฒนา gRPC Service ด้วยภาษา Go โดยมีการเรียกใช้งานทั้งแบบดึงข้อมูลนักศึกษาทีละคน และดึงข้อมูลนักศึกษาหลายคน พร้อมทั้งมีการเพิ่ม field ใหม่ในข้อมูล

---

## 📌 ความสามารถของระบบ (Features)

* ดึงข้อมูลนักศึกษาตาม ID (`GetStudent`)
* ดึงรายการนักศึกษาทั้งหมด (`ListStudents`)
* เพิ่มข้อมูลเบอร์โทรศัพท์ (`phone`) ใน Student

---

## 🏗️ โครงสร้างโปรเจกต์

```text
grpc-student/
│
├── proto/
│   └── student.proto
│
├── studentpb/
│   ├── student.pb.go
│   └── student_grpc.pb.go
│
├── server/
│   └── server.go
│
├── client/
│   └── client.go
│
├── go.mod
└── go.sum
```

---

## ⚙️ สิ่งที่ต้องมี (Requirements)

* ติดตั้ง Go
* ติดตั้ง Protocol Buffers (`protoc`)
* ติดตั้ง Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## 🚀 วิธีการใช้งาน

### 1. Generate Code จาก proto

```bash
protoc --go_out=. --go-grpc_out=. proto/student.proto
```

---

### 2. รัน Server

```bash
go run server/server.go
```

Server จะทำงานที่:

```
localhost:50051
```

---

### 3. รัน Client

```bash
go run client/client.go
```

---

## 📡 รายละเอียด RPC Methods

### 🔹 GetStudent

ใช้สำหรับดึงข้อมูลนักศึกษาตาม ID

---

### 🔹 ListStudents

ใช้สำหรับดึงรายการนักศึกษาทั้งหมด

---

## 🧪 ตัวอย่างผลลัพธ์

```text
Student List:
ID: 101 | Name: Alice Johnson | Major: Computer Science | Email: alice@uni.com | Phone: 0812345678
ID: 102 | Name: Bob Smith | Major: IT | Email: bob@uni.com | Phone: 0898765432

Student Info:
ID: 101
Name: Alice Johnson
Major: Computer Science
Email: alice@university.com
Phone: 0812345678
```

---

## 🧠 สิ่งที่ได้เรียนรู้

* การสร้าง gRPC Service ด้วย Protocol Buffers
* การ generate code จากไฟล์ `.proto`
* การพัฒนา Server และ Client ด้วยภาษา Go
* การเพิ่ม RPC Method ใหม่ในระบบ
* การปรับปรุงโครงสร้างข้อมูล (เพิ่ม field ใหม่)

---

## 📌 งานที่ทำ (Assignment)

### ✅ Task 1

* เพิ่ม `ListStudents` RPC Method
* แก้ไข proto และ generate code ใหม่
* เขียน server ให้รองรับการส่งข้อมูลหลายรายการ
* แก้ client ให้สามารถเรียกและแสดงผล list ได้

### ✅ Task 2

* เพิ่ม field `phone` ใน `StudentResponse`
* แก้ server ให้ส่งข้อมูล phone
* แก้ client ให้แสดงผล phone

---
