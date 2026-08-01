#!/data/data/com.termux/files/usr/bin/bash

# File database & log
DB="users.db"
LOGDIR="logs"
mkdir -p "$LOGDIR"

# Cek dan install gaya
for pkg in lolcat figlet toilet; do
  command -v $pkg >/dev/null 2>&1 || {
    echo "⏳ Menginstall $pkg..."
    pkg install -y $pkg
  }
done

# Fungsi loading
function loading() {
  echo -n "⏳ Proses"
  for i in {1..3}; do
    echo -n "."
    sleep 0.5
  done
  echo
}

# Fungsi update saldo
function update_saldo() {
  local user=$1
  local change=$2
  saldo=$(grep "^$user:" "$DB" | cut -d: -f3)
  new_saldo=$((saldo + change))
  sed -i "s/^$user:[^:]*:[0-9]*/$user:$pass:$new_saldo/" "$DB"
}

# Fungsi dashboard
function dashboard() {
  local username=$1
  while true; do
    clear
    figlet "KONGPAY" | lolcat
    echo -e "\e[1;36mWelcome, $username\e[0m" | lolcat
    echo -e "1. Cek Saldo\n2. Top Up\n3. Transfer\n4. QR Payment\n5. Riwayat\n6. Logout"
    read -p "Pilih menu: " menu

    case $menu in
      1) cek_saldo "$username" ;;
      2) top_up "$username" ;;
      3) transfer "$username" ;;
      4) qr_payment "$username" ;;
      5) cat "$LOGDIR/riwayat-$username.log" 2>/dev/null || echo "Belum ada riwayat." ; read -p "Enter untuk lanjut..." ;;
      6) echo "Logout..."; sleep 1; break ;;
      *) echo "❌ Pilihan tidak valid" ; sleep 1 ;;
    esac
  done
}

# Fungsi cek saldo
function cek_saldo() {
  saldo=$(grep "^$1:" "$DB" | cut -d: -f3)
  echo -e "\e[32m💰 Saldo: Rp $saldo\e[0m"
  sleep 2
}

# Fungsi top up
function top_up() {
  read -p "Jumlah Top Up: " amount
  update_saldo "$1" "$amount"
  echo -e "\e[32m✅ Top up Rp $amount berhasil\e[0m"
  echo "$(date '+%F %T') TOPUP +$amount" >> "$LOGDIR/riwayat-$1.log"
  sleep 2
}

# Fungsi transfer
function transfer() {
  read -p "Tujuan transfer: " target
  read -p "Jumlah: " amount
  saldo=$(grep "^$1:" "$DB" | cut -d: -f3)
  if ! grep -q "^$target:" "$DB"; then
    echo -e "\e[31m❌ User tidak ditemukan\e[0m"
    sleep 2
    return
  fi
  if (( saldo < amount )); then
    echo -e "\e[31m❌ Saldo tidak cukup\e[0m"
    sleep 2
    return
  fi
  update_saldo "$1" "-$amount"
  pass=$(grep "^$target:" "$DB" | cut -d: -f2)
  sed -i "s/^$target:[^:]*:[0-9]*/$target:$pass:$(( $(grep "^$target:" "$DB" | cut -d: -f3) + amount ))/" "$DB"
  echo -e "\e[32m✅ Transfer Rp $amount ke $target berhasil\e[0m"
  echo "$(date '+%F %T') TRANSFER -$amount ke $target" >> "$LOGDIR/riwayat-$1.log"
  echo "$(date '+%F %T') TRANSFER +$amount dari $1" >> "$LOGDIR/riwayat-$target.log"
  sleep 2
}

# Fungsi QR Payment
function qr_payment() {
  read -p "Masukkan Kode Merchant: " merchant
  read -p "Masukkan nominal pembayaran: " amount
  saldo=$(grep "^$1:" "$DB" | cut -d: -f3)
  if (( saldo < amount )); then
    echo -e "\e[31m❌ Saldo tidak cukup!\e[0m"
    sleep 2
    return
  fi
  update_saldo "$1" "-$amount"
  echo -e "\e[32m✅ Pembayaran QR Rp $amount ke $merchant berhasil!\e[0m"
  echo "$(date '+%F %T') QRPAY -$amount ke $merchant" >> "$LOGDIR/riwayat-$1.log"
  sleep 2
}

# Fungsi login
function login() {
  read -p "Username: " user
  read -s -p "Password: " pass
  echo
  if grep -q "^$user:$pass:" "$DB"; then
    dashboard "$user"
  else
    echo -e "\e[31m❌ Login gagal\e[0m"
    sleep 2
  fi
}

# Fungsi daftar
function daftar() {
  read -p "Buat username: " user
  grep -q "^$user:" "$DB" && echo -e "\e[31m❌ Username sudah terdaftar\e[0m" && return
  read -s -p "Buat password: " pass
  echo
  echo "$user:$pass:0" >> "$DB"
  echo -e "\e[32m✅ Registrasi berhasil\e[0m"
  sleep 2
}

# Main Menu
while true; do
  clear
  figlet "KONGPAY" | lolcat
  echo -e "1. Login\n2. Daftar\n3. Keluar"
  read -p "Pilih: " pilih
  case $pilih in
    1) login ;;
    2) daftar ;;
    3) echo "Keluar..."; exit ;;
    *) echo "Pilihan salah"; sleep 1 ;;
  esac
done
