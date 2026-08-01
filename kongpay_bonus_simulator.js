#!/bin/bash

# KONGPAY BONUS FEATURE SIMULATOR
# ----------------------------------
# Tambahan fitur keren untuk KONGPAY: QR Payment Simulasi, Notifikasi Palsu, dan Style Terminal

# Dependencies: lolcat, figlet, toilet, tput

# Cek dependencies
check_dependencies() {
    for cmd in figlet lolcat toilet tput; do
        if ! command -v $cmd &>/dev/null; then
            echo "$cmd belum terinstall. Install dulu ya! (pkg install $cmd)" | lolcat
            exit 1
        fi
    done
}

# Header
print_header() {
    clear
    figlet "KONGPAY" | lolcat
    echo "=====================================" | lolcat
    echo " Simulasi Fitur Tambahan KongPay" | lolcat
    echo "=====================================" | lolcat
}

# Simulasi QR Payment
simulate_qr_payment() {
    echo -e "\n[📷] Menyalakan kamera untuk scan QR..." | lolcat
    sleep 2
    echo -e "[✅] QR Code dikenali: TOKO SOTO BANG JALI" | lolcat
    sleep 1
    echo -e "[💸] Pembayaran sebesar Rp25.000 diproses..." | lolcat
    sleep 2
    echo -e "[🎉] Transaksi berhasil!" | lolcat
    echo -e "[🧾] Saldo Tersisa: Rp75.000" | lolcat
    echo ""
}

# Notifikasi palsu
fake_notifications() {
    echo -e "\n📲 Notifikasi Masuk" | toilet -f term -F border --gay
    echo -e "\n[🔔] Anda menerima Rp1.000.000 dari BOS BESAR" | lolcat
    sleep 2
    echo -e "[🔔] Tagihan listrik berhasil dibayar (Rp150.000)" | lolcat
    sleep 2
    echo -e "[🔔] Cashback Rp50.000 masuk ke dompet KONGPAY Anda!" | lolcat
    echo ""
}

# Menu
main_menu() {
    print_header
    echo -e "1. Simulasi QR Payment"
    echo -e "2. Notifikasi Palsu"
    echo -e "3. Keluar"
    echo -n -e "\nPilih opsi [1-3]: "
    read pilihan
    case $pilihan in
        1)
            simulate_qr_payment
            ;;
        2)
            fake_notifications
            ;;
        3)
            echo -e "\nTerima kasih! Tetap semangat ngoding. 🖤" | lolcat
            exit 0
            ;;
        *)
            echo -e "Pilihan tidak valid!" | lolcat
            ;;
    esac
    echo -e "\nTekan Enter untuk kembali ke menu..."
    read
    main_menu
}

# Jalankan
check_dependencies
main_menu
