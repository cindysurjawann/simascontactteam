import { useState, useEffect } from "react";
import Header from "~/components/Header";
import { useRouter } from "next/router";
import Image from "next/future/image";
import cslaki from "../public/cslaki.png";
import info from "../public/info.png";
import UserFooter from "../components/userfooter";

export default function Nasabah() {
  const [showCS, setShowCS] = useState(false);
  const [showInfo, setShowInfo] = useState(false);
  const [load, setLoading] = useState(true);
  const [loadWA, setLoadingWA] = useState(false);
  const [linkWa, setLinkWA] = useState("i");
  const router = useRouter();

  const kontakresmi = () => {
    router.push("/nasabah/kontakresmi");
  };

  const zoom = () => {
    router.push("/nasabah/zoom");
  };

  const asuransi = () => {
    router.push("/nasabah/asuransi");
  };

  const promo = () => {
    router.push("/nasabah/promo");
  };

  const wa = async () => {
    if (linkWa == "") {
      setLoadingWA(true);
      const res = await getLinkWa();
      if (!res) {
        alert("Gagal ke Wa");
        return;
      }
    }
  };
  const getLinkWa = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}get-link/WA`);
      if (res.status != 200) {
        throw "gagal mendapatkan pesan"();
      }
      const data = await res.json();
      if (!data.data) {
        throw "gagal mendapatkan data"();
      }
      setLinkWA(data.data.linkvalue);
      setLoading(false);
      return true;
    } catch (e) {
      setLoading(false);
      return false;
    }
  };
  useEffect(() => {
    getLinkWa();
  }, []);
  return (
    <div>
      <Header />
      <div className="container pt-5 ">
        <div className="row my-5" id="layanancs">
          <div className="col-6 border-end-0 border-3">
            <Image src={cslaki} width={500} height={400} alt={"info"} />
          </div>
          <div className="col-6">
            <div className="mx-auto">
              <h3 className="fw-bold mb-3">Layanan Customer Service</h3>
              <p style={{ textAlign: "justify" }}>
                Layanan customer service sinarmas hadir untuk menyelesaikan masalah anda terkait layanan yang berada di sinarmas. Layanan CS sinarmas tidak hanya menangani permasalahan untuk nasabah, tetapi terbuka untuk nasabah yang mau
                menjelajahi kebebasan finansial bersama sinarmas.
              </p>
            </div>

            <div className="mt-5">
              <p>Anda punya pertanyaan?</p>
              <button
                className="btn fw-bold"
                onClick={() => setShowCS(!showCS)}
                style={{
                  background: "#CC100F",
                  color: "white",
                  height: "50px",
                }}
              >
                Tanya Sekarang
              </button>
              {showCS ? (
                load ? (
                  <div className="mt-3">
                    <p>Please wait</p>
                  </div>
                ) : (
                  <div className="d-flex mt-3">
                    <button className="btn me-2" style={{ background: "#CC100F", color: "white" }} onClick={zoom}>
                      Video Zoom
                    </button>
                    <a className="btn" style={{ background: "#CC100F", color: "white" }} onClick={wa} href={linkWa}>
                      {loadWA ? "Please wait" : ""} Whatsapp
                    </a>
                  </div>
                )
              ) : (
                ""
              )}
            </div>
          </div>
        </div>
        <br />
        <hr />
        <br />
        <div className="row my-5" id="layananinfo">
          <div className="col-6 border-end-0 border-3">
            <Image src={info} width={500} height={400} style={{ borderRadius: "75%" }} alt={"info"} />
          </div>
          <div className="col-6">
            <div className="mx-auto">
              <h3 className="fw-bold mb-3">Pusat Informasi</h3>
              <p style={{ textAlign: "justify" }}>Pusat Informasi Bank Sinarmas hadir untuk menyediakan informasi-informasi terkini mengenai Promosi, Asuransi, dan Kontak Resmi.</p>
            </div>

            <div className="mt-5">
              <p>Anda Mencari Informasi?</p>
              <button
                className="btn fw-bold"
                onClick={() => setShowInfo(!showInfo)}
                style={{
                  background: "#CC100F",
                  color: "white",
                  height: "50px",
                }}
              >
                Eksplor Sekarang
              </button>
              {showInfo ? (
                load ? (
                  <div className="mt-3">
                    <p>Please wait</p>
                  </div>
                ) : (
                  <div className="d-flex mt-3">
                    <button className="btn me-2" style={{ background: "#CC100F", color: "white" }} onClick={asuransi}>
                      Asuransi
                    </button>
                    <button className="btn me-2" style={{ background: "#CC100F", color: "white" }} onClick={promo}>
                      Promo
                    </button>
                    <button className="btn" style={{ background: "#CC100F", color: "white" }} onClick={kontakresmi}>
                      Kontak Resmi
                    </button>
                  </div>
                )
              ) : (
                ""
              )}
            </div>
          </div>
        </div>
      </div>
      <UserFooter />
    </div>
  );
}
