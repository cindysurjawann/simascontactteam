import { useState, useEffect } from "react";
import Header from "~/components/Header";
import { useRouter } from "next/router";
import Image from "next/future/image";
import desktopzoom from "../../public/desktopzoom.png";
import UserFooter from "../../components/userfooter";
import styles from "./Zoom.module.scss";

export default function Zoom() {
  const [link, setLink] = useState(false);
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  const [data, setData] = useState('');

  const getLocation = async (lat, lang) => {
    try {
      const newUrl = `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${lat}&longitude=${lang}&localityLanguage=id`
      const res = await fetch(newUrl);
      const data = await res.json();
      setData(data);
      console.log(data);

    }
    catch (error) {
      alert("Gagal get location");
    }
  }

  const getLinkZoom = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}get-link/Zoom`);
      if (res.status != 200) {
        throw "gagal mendapatkan pesan"();
      }
      const data = await res.json();
      if (!data.data) {
        throw "gagal mendapatkan data"();
      }
      setLink(data.data.linkvalue);
      setLoading(false);
      return true;
    } catch (e) {
      if (typeof e === "string") {
        alert("Link Zoom tidak ada, silahkan merefresh ulang");
      }
      return false;
    }
  };

  const postDataZoom = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      nama: formData.get("namaZoom"),
      email: formData.get("emailZoom"),
      kategori: formData.get("kategoriZoom"),
      keterangan: formData.get("keluhanZoom"),
      lokasi: data !== null ? data.city : "",
    };
    console.log(body);
    try {
      await fetch(`${process.env.NEXT_PUBLIC_URL}createzoomhistory`, {
        method: "POST",
        body: JSON.stringify(body),
      });
      if (link == "") {
        alert("link zoom sedang bermasalah, silahkan merefresh ulang");
        return;
      }
      router.push(link);
    } catch (e) {
      if (typeof e === "string") {
        alert("Gagal menginputkan form data diri, silahkan refresh ulang");
      }
      return false;
    }
  };

  useEffect(() => {
    getLinkZoom();
    navigator.geolocation.getCurrentPosition(function(position) {
        getLocation(position.coords.latitude, position.coords.longitude)
    });
  }, []);

  return (
    <div>
      <Header />
      <div className="col-6 border-end-0 border-3"></div>
      <div className="container-fluid mt-5  ">
        <h2 className="ms-3 fw-bold">Layanan CS - Video Zoom</h2>
        <div className="row mt-4">
          <div className="col-6">
            <Image
              src={desktopzoom}
              width={500}
              height={400}
              alt={"desktopzoom"}
              priority={true}
            />
          </div>
          <div className="col-6">
            <h3 className="fw-bold">Proses Video Zoom</h3>
            <div style={{ maxWidth: "80%" }} className="mt-4">
              <ol>
                <li>
                  Silahkan join room zoom menggunakan nama dan email yang sesuai
                  diisikan di data diri, apabila tidak sesuai maka tidak akan
                  diproses{" "}
                </li>
                <li>
                  Tunggu terlebih dahulu di ruang utama hingga customer service
                  mengundangan ke breakout room
                </li>
                <li>Silahkan join breakout room </li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <br />
      <hr />
      <br />
      <div className={styles.container}>
        <form id="formZoom" onSubmit={postDataZoom}>
          <h3 className={styles.titleForm}>Isi Data Diri</h3>

          <label htmlFor="namaZoom" className={styles.label}>
            Nama
          </label>
          <input
            type="text"
            placeholder="Masukan Nama"
            id="namaZoom"
            name="namaZoom"
            className={styles.input}
            required
          />

          <label htmlFor="emailZoom" className={styles.label}>
            Email
          </label>
          <input
            type="email"
            placeholder="Masukkan Email"
            id="emailZoom"
            name="emailZoom"
            className={styles.input}
            required
          />

          <label htmlFor="kategoriZoom" className={styles.label}>
            Kategori
          </label>
          <select
            id="kategoriZoom"
            name="kategoriZoom"
            className={styles.input}
          >
            <option value="perbankan">Perbankan</option>
            <option value="kartuKredit">Kartu Kredit</option>
            <option value="digitalLoan">Digital Loan</option>
            <option value="merchantQR">Merchant QRIS</option>
            <option value="pin">Mengganti PIN Channel</option>
            <option value="custcare">Berbicara dengan CustCare</option>
          </select>

          <label htmlFor="keluhanZoom" className={styles.label}>
            Keluhan
          </label>
          <textarea
            id="keluhanZoom"
            name="keluhanZoom"
            placeholder="Masukkan Keluhan Anda"
            className={styles.input}
            required
          />

          {loading ? (
            <div>Please Wait</div>
          ) : (
            <button className={styles.button}>Melanjutkan ke Zoom</button>
          )}
        </form>
      </div>
      <UserFooter />
    </div>
  );
}
