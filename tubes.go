package main

import (
    "fmt"
    "os"
    "time"
)

type User struct {
    Username string
    Password string
}

type Barang struct {
    ID    int
    Name  string
    Stock int
    Harga int
}

type Inventori struct {
    Barangs    [100]Barang
    Transaksis [100]Transaksi
    BarangCount    int
    TransaksiCount int
}

type Transaksi struct {
    Time     time.Time
    ItemID   int
    ItemName string
    Tipe     string
    Jumlah   int
    Editor   string
}

var loggedInUser *User

func main() {
    registeredUsers := [100]User{
        {Username: "admin", Password: "admin"},
    }
    userCount := 1

    inventori := Inventori{
        Barangs: [100]Barang{
            {ID: 1, Name: "aqua", Stock: 10, Harga: 5000},
            {ID: 2, Name: "le minerale", Stock: 20, Harga: 10000},
            {ID: 3, Name: "beras", Stock: 15, Harga: 7500},
            {ID: 4, Name: "gula", Stock: 5, Harga: 2000},
            {ID: 5, Name: "masako", Stock: 8, Harga: 3000},
        },
        BarangCount: 5,
    }

    for {
        if loggedInUser == nil {
            fmt.Println("\n===== Selamat Datang =====")
            fmt.Println("1. Login")
            fmt.Println("2. Register")
            fmt.Println("3. Keluar")
            fmt.Print("Pilihan Anda: ")

            var loginChoice int
            fmt.Scan(&loginChoice)

            switch loginChoice {
            case 1:
                loggedInUser = Login(registeredUsers[:userCount])
            case 2:
                registeredUser := Register(&registeredUsers, &userCount)
                if registeredUser.Username != "" {
                    registeredUsers[userCount] = registeredUser
                    userCount++
                    fmt.Println("Silakan login.")
                }
            case 3:
                fmt.Println("Terima kasih!")
                os.Exit(0)
            default:
                fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
            }
        } else {
            fmt.Printf("\nHalo %v", loggedInUser.Username)
            fmt.Println("\n===== Menu =====")
            fmt.Println("1. Tambah Item")
            fmt.Println("2. Edit Item")
            fmt.Println("3. Hapus Item")
            fmt.Println("4. Tampilkan Inventori")
            fmt.Println("5. Tampilkan Inventori Berurut")
            fmt.Println("6. Tampilkan Catatan Transaksi")
            fmt.Println("7. Cari barang berdasarkan ID atau Nama Barang")
            fmt.Println("8. Transaksi Pembayaran")
            fmt.Println("9. Keluar")
            fmt.Print("Pilihan Anda: ")

            var choice int
            fmt.Scan(&choice)

            switch choice {
            case 1:
                Tambah(&inventori)
            case 2:
                Edit(&inventori)
            case 3:
                Hapus(&inventori)
            case 4:
                Tampilkan(&inventori)
            case 5:
                PilihPengurutan(&inventori)
            case 6:
                TampilkanCatatan(&inventori)
            case 7:
                CariBarang(&inventori)
            case 8:
                TransaksiPembayaran(&inventori)
            case 9:
                fmt.Println("Terima kasih!")
                os.Exit(0)
            default:
                fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
            }
        }
    }
}

func Login(registeredUsers []User) *User {
    var username, password string
    fmt.Println("\n===== Login =====")
    fmt.Print("Masukkan username: ")
    fmt.Scan(&username)
    fmt.Print("Masukkan password: ")
    fmt.Scan(&password)

    for i := 0; i < len(registeredUsers); i++ {
        if registeredUsers[i].Username == username && registeredUsers[i].Password == password {
            fmt.Println("Login berhasil!")
            return &registeredUsers[i]
        }
    }
    fmt.Println("Username atau password salah.")
    return nil
}

func Register(registeredUsers *[100]User, userCount *int) User {
    var username, password string
    fmt.Println("\n===== Registrasi =====")
    fmt.Print("Masukkan username: ")
    fmt.Scan(&username)

    for i := 0; i < *userCount; i++ {
        user := (*registeredUsers)[i]
        if user.Username == username {
            fmt.Println("Username sudah terdaftar. Silakan gunakan username lain.")
            return User{Username: "", Password: ""}
        }
    }

    fmt.Print("Masukkan password: ")
    fmt.Scan(&password)

    fmt.Println("Registrasi berhasil!")
    return User{Username: username, Password: password}
}

func Tambah(inv *Inventori) {
    if inv.BarangCount >= 100 {
        fmt.Println("Inventori penuh.")
        return
    }

    fmt.Println("\n===== Tambah Item =====")

    var name string
    fmt.Print("Nama: ")
    fmt.Scan(&name)

    var stock int
    fmt.Print("Stok: ")
    fmt.Scan(&stock)

    var harga int
    fmt.Print("Harga: ")
    fmt.Scan(&harga)

    itemID := 1
    if inv.BarangCount > 0 {
        itemID = inv.Barangs[inv.BarangCount-1].ID + 1
    }

    item := Barang{ID: itemID, Name: name, Stock: stock, Harga: harga}
    inv.Barangs[inv.BarangCount] = item
    inv.BarangCount++

    Catat(inv, itemID, name, "Masuk", stock)

    fmt.Println("Item berhasil ditambahkan.")
}

func Catat(inv *Inventori, itemID int, itemName string, transactionType string, amount int) {
    if inv.TransaksiCount >= 100 {
        fmt.Println("Catatan transaksi penuh.")
        return
    }

    transaction := Transaksi{
        Time:     time.Now(),
        ItemID:   itemID,
        ItemName: itemName,
        Tipe:     transactionType,
        Jumlah:   amount,
        Editor:   loggedInUser.Username,
    }
    inv.Transaksis[inv.TransaksiCount] = transaction
    inv.TransaksiCount++

    if transactionType == "Keluar" {
        fmt.Printf("Stok %s berkurang sebanyak %d.\n", itemName, amount)
    }
}

func Edit(inv *Inventori) {
    fmt.Println("\n===== Edit Item =====")

    var id int
    fmt.Print("ID Item yang akan diedit: ")
    fmt.Scan(&id)

    index := CariIndeks(inv, id)
    if index == -1 {
        fmt.Println("Item dengan ID tersebut tidak ditemukan.")
        return
    }

    item := inv.Barangs[index]

    var pil int
    fmt.Println("Nama barang: ", item.Name)
    fmt.Println("Stok barang: ", item.Stock)
    fmt.Println("Harga barang: ", item.Harga)
    fmt.Print("\n1. Edit nama\n2. Edit stok\n3. Edit harga\n")
    fmt.Print("Pilih: ")
    fmt.Scan(&pil)

    switch pil {
    case 1:
        var name string
        fmt.Print("Nama Baru: ")
        fmt.Scan(&name)
        inv.Barangs[index].Name = name
        Catat(inv, id, name, "EDIT NAMA", item.Stock)
        fmt.Println("Item berhasil diubah.")
    case 2:
        var stock int
        fmt.Print("Stok Baru: ")
        fmt.Scan(&stock)
        amountDifference := stock - item.Stock
        inv.Barangs[index].Stock = stock
        var transactionType string
        if amountDifference > 0 {
            transactionType = "Stok Bertambah"
        } else if amountDifference < 0 {
            transactionType = "Stok Berkurang"
            amountDifference = -amountDifference
        } else {
            transactionType = "Stok Tidak Berubah"
        }
        Catat(inv, id, item.Name, transactionType, amountDifference)
        fmt.Println("Item berhasil diubah.")
    case 3:
        var harga int
        fmt.Print("Harga baru: ")
        fmt.Scan(&harga)
        inv.Barangs[index].Harga = harga
        Catat(inv, id, item.Name, "EDIT HARGA", item.Stock)
        fmt.Println("Item berhasil diubah.")
    default:
        fmt.Print("Pilihan tidak valid.")
    }
}

func Hapus(inv *Inventori) {
    fmt.Println("\n===== Hapus Item =====")

    var id int
    fmt.Print("ID Item yang akan dihapus: ")
    fmt.Scan(&id)

    index := CariIndeks(inv, id)
    if index == -1 {
        fmt.Println("Item dengan ID tersebut tidak ditemukan.")
        return
    }

    item := inv.Barangs[index]

    // Geser elemen-elemen setelah item yang dihapus ke kiri
    for i := index; i < inv.BarangCount-1; i++ {
        inv.Barangs[i] = inv.Barangs[i+1]
    }
    // Kurangi jumlah barang
    inv.BarangCount--

    Catat(inv, item.ID, item.Name, "Keluar", item.Stock)

    fmt.Println("Item berhasil dihapus.")
}

func TransaksiPembayaran(inv *Inventori) {
    fmt.Println("\n===== Transaksi Pembayaran =====")
    fmt.Println("1. Cari berdasarkan ID")
    fmt.Println("2. Cari berdasarkan Nama")
    fmt.Print("Pilihan Anda: ")

    var pilihan int
    fmt.Scan(&pilihan)

    var index int
    switch pilihan {
    case 1:
        var id int
        fmt.Print("Masukkan ID Barang: ")
        fmt.Scan(&id)
        index = CariIndeks(inv, id)
        if index == -1 {
            fmt.Println("Item dengan ID tersebut tidak ditemukan.")
            return
        }
    case 2:
        var name string
        fmt.Print("Masukkan Nama Barang: ")
        fmt.Scan(&name)
        index = -1
        for i := 0; i < len(inv.Barangs); i++ {
            if inv.Barangs[i].Name == name {
                index = i
                break
            }
        }
        if index == -1 {
            fmt.Println("Item dengan nama tersebut tidak ditemukan.")
            return
        }
    default:
        fmt.Println("Pilihan tidak valid.")
        return
    }

    var jumlah int
    fmt.Print("Masukkan Jumlah yang terbeli: ")
    fmt.Scan(&jumlah)

    if inv.Barangs[index].Stock < jumlah {
        fmt.Println("Stok tidak mencukupi.")
        return
    }

    stokSebelum := inv.Barangs[index].Stock
    inv.Barangs[index].Stock -= jumlah
    Catat(inv, inv.Barangs[index].ID, inv.Barangs[index].Name, "Keluar", jumlah)
    fmt.Printf("Transaksi berhasil. Stok %s berkurang dari %d menjadi %d.\n", inv.Barangs[index].Name, stokSebelum, inv.Barangs[index].Stock)
}


func PilihPengurutan(inv *Inventori) {
    fmt.Println("\n===== Pilih Pengurutan =====")
    fmt.Println("1. Berdasarkan Stok")
    fmt.Println("2. Berdasarkan Harga")
    fmt.Print("Pilihan Anda: ")

    var pilihan int
    fmt.Scan(&pilihan)

    fmt.Println("\n===== Pilih Urutan =====")
    fmt.Println("1. Ascending")
    fmt.Println("2. Descending")
    fmt.Print("Pilihan Anda: ")

    var urutan int
    fmt.Scan(&urutan)

    switch pilihan {
    case 1:
        if urutan == 1 {
            TampilkanBerurutStokAscending(inv)
        } else {
            TampilkanBerurutStokDescending(inv)
        }
    case 2:
        if urutan == 1 {
            TampilkanBerurutHargaAscending(inv)
        } else {
            TampilkanBerurutHargaDescending(inv)
        }
    default:
        fmt.Println("Pilihan tidak valid.")
    }
}

func TampilkanBerurutStokAscending(inv *Inventori) {
    fmt.Println("\n===== Inventori (Stok Terkecil ke Terbesar) =====")
    if inv.BarangCount == 0 {
        fmt.Println("Inventori kosong.")
        return
    }

    n := inv.BarangCount
    // selection sort dimulai
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if inv.Barangs[j].Stock < inv.Barangs[minIndex].Stock {
                minIndex = j
            }
        }
        temp := inv.Barangs[i]
        inv.Barangs[i] = inv.Barangs[minIndex]
        inv.Barangs[minIndex] = temp
    }

    for i := 0; i < n; i++ {
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%d Stok: %d\n", inv.Barangs[i].ID, inv.Barangs[i].Name, inv.Barangs[i].Harga, inv.Barangs[i].Stock)
    }
}

func TampilkanBerurutStokDescending(inv *Inventori) {
    fmt.Println("\n===== Inventori (Stok Terbesar ke Terkecil) =====")
    if inv.BarangCount == 0 {
        fmt.Println("Inventori kosong.")
        return
    }

    n := inv.BarangCount
    // insertion sort dimulai
    for i := 1; i < n; i++ {
        key := inv.Barangs[i]
        j := i - 1
        for j >= 0 && inv.Barangs[j].Stock < key.Stock {
            inv.Barangs[j+1] = inv.Barangs[j]
            j = j - 1
        }
        inv.Barangs[j+1] = key
    }

    for i := 0; i < n; i++ {
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%d Stok: %d\n", inv.Barangs[i].ID, inv.Barangs[i].Name, inv.Barangs[i].Harga, inv.Barangs[i].Stock)
    }
}

func TampilkanBerurutHargaAscending(inv *Inventori) {
    fmt.Println("\n===== Inventori (Harga Terkecil ke Terbesar) =====")
    if inv.BarangCount == 0 {
        fmt.Println("Inventori kosong.")
        return
    }

    n := inv.BarangCount
    // Selection sort
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if inv.Barangs[j].Harga < inv.Barangs[minIndex].Harga {
                minIndex = j
            }
        }
        // Tukar elemen
        temp := inv.Barangs[minIndex]
        inv.Barangs[minIndex] = inv.Barangs[i]
        inv.Barangs[i] = temp
    }

    for i := 0; i < n; i++ {
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%d, Stok: %d\n", inv.Barangs[i].ID, inv.Barangs[i].Name, inv.Barangs[i].Harga, inv.Barangs[i].Stock)
    }
}

func TampilkanBerurutHargaDescending(inv *Inventori) {
    fmt.Println("\n===== Inventori (Harga Terbesar ke Terkecil) =====")
    if inv.BarangCount == 0 {
        fmt.Println("Inventori kosong.")
        return
    }

    n := inv.BarangCount
    // insertion sort dimulai
    for i := 1; i < n; i++ {
        key := inv.Barangs[i]
        j := i - 1
        for j >= 0 && inv.Barangs[j].Harga < key.Harga {
            inv.Barangs[j+1] = inv.Barangs[j]
            j = j - 1
        }
        inv.Barangs[j+1] = key
    }

    for i := 0; i < n; i++ {
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%d Stok: %d\n", inv.Barangs[i].ID, inv.Barangs[i].Name, inv.Barangs[i].Harga, inv.Barangs[i].Stock)
    }
}

func CariBarang(inv *Inventori) {
    fmt.Println("\n===== Cari Barang =====")
    fmt.Println("1. Cari berdasarkan ID")
    fmt.Println("2. Cari berdasarkan Nama")
    fmt.Print("Pilihan Anda: ")

    var pilihan int
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
        var id int
        fmt.Print("Masukkan ID Barang: ")
        fmt.Scan(&id)
        index := CariIndeks(inv, id)
        if index != -1 {
            item := inv.Barangs[index]
            fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: Rp%d\n", item.ID, item.Name, item.Stock, item.Harga)
        } else {
            fmt.Println("Barang tidak ditemukan.")
        }
    case 2:
        var name string
        fmt.Print("Masukkan Nama Barang: ")
        fmt.Scan(&name)
        CariNama(inv, name)
    default:
        fmt.Println("Pilihan tidak valid.")
    }
}

// implementasi sequential search
func CariNama(inv *Inventori, name string) {
    found := false
    for i := 0; i < inv.BarangCount; i++ {
        item := inv.Barangs[i]
        if item.Name == name {
            fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: Rp%d\n", item.ID, item.Name, item.Stock, item.Harga)
            found = true
        }
    }
    if !found {
        fmt.Println("Barang tidak ditemukan.")
    }
}

// implementasi binary search
func CariIndeks(inv *Inventori, id int) int {
    low := 0
    high := inv.BarangCount - 1

    for low <= high {
        mid := (low + high) / 2
        if inv.Barangs[mid].ID == id {
            return mid
        } else if inv.Barangs[mid].ID < id {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

// tampilkan catatan inventori
func Tampilkan(inv *Inventori) {
    fmt.Println("\n===== Inventori =====")
    if inv.BarangCount == 0 {
        fmt.Println("Inventori kosong.")
        return
    }
    for i := 0; i < inv.BarangCount; i++ {
        item := inv.Barangs[i]
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%d Stok: %d\n", item.ID, item.Name, item.Harga, item.Stock)
    }
}

// tampilkan catatan transaksi
func TampilkanCatatan(inv *Inventori) {
    fmt.Println("\n===== Catatan Transaksi =====")
    if inv.TransaksiCount == 0 {
        fmt.Println("Belum ada transaksi.")
        return
    }
    for i := 0; i < inv.TransaksiCount; i++ {
        transaction := inv.Transaksis[i]
        fmt.Printf("%d. Time: %s, Nama Barang: %s, Tipe transaksi: %s, Jumlah: %d, Editor: %s\n",
            i+1, transaction.Time.Format("2006-01-02 15:04:05"), transaction.ItemName, transaction.Tipe, transaction.Jumlah, transaction.Editor)
    }
}