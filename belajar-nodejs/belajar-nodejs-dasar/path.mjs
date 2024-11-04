import path from "path";

const file = "/Users/luffy/hello-world.txt";

// Mengambil nama file dari path.
console.info(path.basename(file)); // Output: "hello-world.txt"

// Separator untuk beberapa path, ";" di Windows dan ":" di POSIX.
console.info(path.delimiter); // Output: ";" atau ":", tergantung OS

// Mengembalikan nama direktori dari path.
console.info(path.dirname(file)); // Output: "/Users/luffy"

// Mengambil ekstensi file.
console.info(path.extname(file)); // Output: ".txt"

// Membangun kembali path dari objek path.
const pathObject = path.parse(file);
console.info(path.format(pathObject)); // Output: "/Users/luffy/hello-world.txt"

// Memeriksa apakah path adalah path absolut.
console.info(path.isAbsolute(file)); // Output: true

// Menggabungkan beberapa path.
console.info(path.join("/Users", "luffy", "projects")); // Output: "/Users/luffy/projects"

// Menormalkan path (menghilangkan elemen redundan).
console.info(path.normalize("/Users/luffy/../luffy/projects")); // Output: "/Users/luffy/projects"

// Menguraikan path menjadi objek.
console.info(path.parse(file));
/* Output:
{
  root: "/",
  dir: "/Users/luffy",
  base: "hello-world.txt",
  ext: ".txt",
  name: "hello-world"
}
*/

// Versi POSIX dari modul path.
console.info(path.posix.join("/Users", "luffy", "projects")); // Output: "/Users/luffy/projects"

// Menghitung path relatif dari satu direktori ke direktori lain.
console.info(path.relative("/Users/luffy", "/Users/luffy/projects")); // Output: "projects"

// Menyelesaikan urutan path menjadi path absolut.
console.info(path.resolve("docs", "file.txt")); // Output: "/current/dir/docs/file.txt"

// Separator untuk memisahkan direktori, "\" di Windows dan "/" di POSIX.
console.info(path.sep); // Output: "/" atau "\", tergantung OS

// Mengonversi path ke bentuk Namespaced di Windows.
console.info(path.toNamespacedPath(file)); // Output khusus di Windows

// Versi Windows dari modul path.
console.info(path.win32.join("C:", "Users", "luffy")); // Output: "C:\Users\luffy" di Windows
