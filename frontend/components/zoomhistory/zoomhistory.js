import style from "../managecs/managecs.module.scss";
import { useState, useEffect } from "react";

const ManageCS = () => {
  const [data, setData] = useState(null);
  
  const getUser = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}getzoomhistory`, {
        method: "GET",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
      });
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

  const column = ["#", "Nama", "Email", "Kategori", "Keterangan", "Lokasi"];

  useEffect(() => {
    getUser();
  }, []);

  return (
    <div className={style.body}>
      <h3>Riwayat Akses Menu Zoom</h3>
      <br />
      <br />
      <table className="table table-hover table-striped">
        <thead>
          <tr>
            {column?.map((item, index) => (
              <th scope="col" className="text-light bg-dark" key={index}>
                {item}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data?.data?.map((item, index) => (
            <tr key={index}>
              <th scope="row">{index + 1}</th>
              <td>{item.nama}</td>
              <td>{item.email}</td>
              <td>{item.kategori}</td>
              <td>{item.keterangan}</td>
              <td>{item.lokasi}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ManageCS;