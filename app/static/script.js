document.getElementById("form-absensi").addEventListener("submit", function(event) {
    event.preventDefault(); // Mencegah reload halaman
    const form = document.getElementById("form-absensi");
    // Ambil nilai dari form
    const name = document.getElementById("name").value;
    const nimValue = document.getElementById("nim").value;
    const mataKuliah = document.getElementById("mata_kuliah").value;
    const jurusan = document.getElementById("jurusan").value;

    // 🔹 Pastikan nim dikonversi ke int
    const nim = parseInt(nimValue, 10);
    if (isNaN(nim) || nim <= 0) {
        alert("❌ NIM harus berupa angka!");
        return;
    }

    const formData = {
        name: name,
        nim: nim, // 🔥 Pastikan ini angka
        mata_kuliah: mataKuliah,
        jurusan: jurusan
    };

    console.log("📡 Data dikirim:", formData); // 🔍 Debugging

    fetch("http://localhost:3000/", {
        method: "POST",
        headers: { "Content-Type": "application/json" }, // 🔥 Wajib JSON
        body: JSON.stringify(formData) // 🔥 Konversi ke JSON
    })
    .then(response => response.json().then(data => ({ status: response.status, body: data })))
    .then(({ status, body }) => {
        if (status === 200) {
            alert("✅ Absensi berhasil dikirim!");
             setTimeout(() => {
                form.reset();
            }, 100);
        } else {
            alert(`❌ Gagal mengirim: ${body.status}`);
        }
    })
    .catch(error => {
        alert("❌ Terjadi kesalahan!");
        console.error("Error:", error);
    });
});
