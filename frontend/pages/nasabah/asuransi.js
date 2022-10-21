import React,{ useState, useEffect } from "react";
import Header from "../../components/Header";
import UserFooter from "../../components/userfooter";
import style from "./asuransi.module.scss";
import Image from "next/future/image";
import jumbotron from "../../public/jumbotron.png";
import ConfirmationModal from "../../components/modals/modalDetailCart";
import { useRouter } from "next/router";

const Promo = () => {
  const [data, setData] = useState(null);
  const [modalOpen, setModalOpen] = React.useState(false);
  const [body, setBodyData] = React.useState("");
  const router = useRouter();
  const format = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
  });
  const promo = () => {
    router.push("/nasabah/promo");
  };

  useEffect(() => {
    getData();
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}getasuransi`);
      if (res.status != 200) {
        throw "gagal mendapatkan pesan"();
      }
      const data = await res.json();
      if (!data.data) {
        throw "gagal mendapatkan data"();
      }
      setData(data);
    } catch (e) {
      alert("Gagal mengambil data");
    }
  };

  return (
    <div>
      <Header />
      <div>
        <Image className={style.jumbotron} src={jumbotron} alt="jumbotron" />
      </div>
      <div className={style.buttonpa}>
        <div>
          <button className={style.buttonpromo }onClick={promo}>Promo</button>
        </div>
        <div>
          <button className={style.buttonasuransiActive}>
            Asuransi
          </button>
        </div>
      </div>
      {getData}
      <div
        className="row justify-content-start"
        style={{ paddingLeft: 80, paddingRight: 80 }}
      >
        {data?.data?.map((item, index) => (
          <div
            key={index}
            className="col-4"
            style={{ paddingLeft: 50, paddingRight: 50 }}
          >
            <div className={style.detailContent}>
              <img
                className={style.imageDummy}
                src={item.foto}
                alt={item.judul}
                width={100}
                height={100}
              />
              <h3 className={style.textContent}>{item.judul}</h3>
              <h5 className={style.textContent} style={{ fontSize: "1rem" }}>
                Premi Per Tahun: {format.format(item.premi)}
              </h5>
              <h5 className={style.textContent} style={{ fontSize: "1rem" }}>
                UP: {format.format(item.uangpertanggungan)}
              </h5>
              <button
                className={style.buttonDetail}
                onClick={() => {
                  setBodyData(item);
                  setModalOpen(true);
                }}
              >
                Lihat Detail
              </button>
            </div>
          </div>
        ))}
      </div>
      {body ? (
        <ConfirmationModal
          show={modalOpen}
          close={() => setModalOpen(false)}
          data={body}
        />
      ) : (
        ""
      )}
      <UserFooter />
    </div>
  );
};

export default Promo;
