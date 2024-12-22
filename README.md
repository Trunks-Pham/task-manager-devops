# Task Manager

Dự án quản lý công việc (Task Manager) với backend viết bằng Golang và frontend bằng React. Dự án này sử dụng Docker để đóng gói và chạy ứng dụng.

## Cấu trúc thư mục

```markdown
task-manager/
├── backend/            # Backend Golang
│   ├── main.go
│   ├── handlers/
│   │   ├── task_handler.go
│   ├── models/
│   │   ├── task.go
│   ├── database/
│   │   ├── db.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
├── frontend/           # Frontend React
│   ├── src/
│   │   ├── App.jsx
│   │   ├── index.jsx
│   │   ├── components/
│   │   │   ├── TaskList.jsx
│   │   │   ├── TaskForm.jsx
│   ├── public/
│   ├── Dockerfile
│   ├── package.json
│   ├── vite.config.js
├── docker-compose.yml
└── README.md
```

## Mô tả

- **Backend (Golang)**: Cung cấp API để quản lý các công việc (tasks), bao gồm các chức năng tạo, xem, cập nhật và xóa công việc.
- **Frontend (React)**: Giao diện người dùng cho phép người dùng quản lý các công việc của mình thông qua các form và danh sách công việc.
- **Docker**: Dự án sử dụng Docker để đóng gói cả backend và frontend vào các container và dễ dàng triển khai trên các môi trường khác nhau.

## Cài đặt và Sử dụng

### Prerequisites

- **Docker**: Bạn cần cài đặt Docker trên máy tính của mình.
- **Docker Compose**: Cài đặt Docker Compose để quản lý các dịch vụ Docker dễ dàng hơn.

### 1. Cài đặt Docker

Trước khi bắt đầu, bạn cần cài đặt Docker và Docker Compose trên máy tính của mình. Vui lòng tham khảo tài liệu chính thức của Docker để cài đặt:

- [Hướng dẫn cài đặt Docker](https://docs.docker.com/get-docker/)
- [Hướng dẫn cài đặt Docker Compose](https://docs.docker.com/compose/install/)

### 2. Build và Chạy Ứng Dụng

Clone dự án từ GitHub (nếu bạn chưa có dự án trên máy):

```bash
git clone https://github.com/Trunks-Pham/task-manager-devops.git
cd task-manager
```

#### Build và khởi động các container Docker

- **Bước 1**: Trong thư mục gốc của dự án, sử dụng Docker Compose để build và chạy các dịch vụ backend và frontend:

```bash
docker-compose up --build
```

Lệnh trên sẽ:
- Build và chạy container backend Golang.
- Build và chạy container frontend React.
- Kết nối các dịch vụ với nhau thông qua Docker Compose.

- **Bước 2**: Truy cập ứng dụng của bạn thông qua trình duyệt:
  - Backend (API): `http://localhost:8080`
  - Frontend (Giao diện người dùng): `http://localhost:3000`

### 3. Cấu hình Dockerfile

#### Backend Dockerfile

Dockerfile trong thư mục `backend/` dùng để build và chạy ứng dụng Golang.

```dockerfile
# Sử dụng image Golang chính thức
FROM golang:1.20

# Đặt thư mục làm việc
WORKDIR /app

# Sao chép go.mod và go.sum và tải dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Sao chép mã nguồn vào container
COPY . ./

# Biên dịch mã nguồn thành file thực thi main
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]
```

#### Frontend Dockerfile

Dockerfile trong thư mục `frontend/` dùng để build và chạy ứng dụng React.

```dockerfile
# Sử dụng image Node.js chính thức
FROM node:16

# Đặt thư mục làm việc
WORKDIR /app

# Sao chép package.json và package-lock.json
COPY package.json ./
COPY package-lock.json ./

# Cài đặt các dependencies
RUN npm install

# Sao chép mã nguồn vào container
COPY . ./

# Build ứng dụng React
RUN npm run build

# Cung cấp dịch vụ trên port 80
EXPOSE 80

# Chạy ứng dụng
CMD ["npm", "start"]
```

### 4. Các API

API trong backend cung cấp các chức năng CRUD cho công việc. Dưới đây là các endpoint chính:

- **GET /tasks**: Lấy danh sách tất cả các công việc.
- **POST /tasks**: Thêm một công việc mới.
- **PUT /tasks/{id}**: Cập nhật thông tin của công việc.
- **DELETE /tasks/{id}**: Xóa công việc theo ID.

### 5. Frontend (React)

Giao diện frontend có các thành phần sau:

- **TaskList**: Hiển thị danh sách các công việc.
- **TaskForm**: Form thêm mới hoặc cập nhật công việc.

### 6. Docker Compose

Dự án sử dụng **docker-compose.yml** để quản lý các container Docker cho backend và frontend.

```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=yourpassword
      - DB_NAME=taskdb
    restart: always

  frontend:
    build: ./frontend
    ports:
      - "3000:80"
    depends_on:
      - backend
    restart: always
```

### 7. Xử lý lỗi

Nếu gặp lỗi khi chạy Docker Compose, kiểm tra các vấn đề sau:
- Kiểm tra Dockerfile và cấu hình đúng đường dẫn cho file thực thi.
- Kiểm tra các container có đang chạy không bằng lệnh `docker ps`.
- Kiểm tra log của container backend: `docker logs <container_id>`.

### 8. Đóng gói và Triển khai

Sau khi hoàn tất, bạn có thể triển khai ứng dụng lên môi trường sản xuất hoặc hosting với Docker. Các bước triển khai sẽ tùy thuộc vào nhà cung cấp dịch vụ như AWS, DigitalOcean, hoặc Heroku.

---

## License

MIT License - Xem [LICENSE](LICENSE) để biết thêm chi tiết.

```