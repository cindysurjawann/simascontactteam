import style from "./managecs.module.scss";
import CreateIcon from "@mui/icons-material/Create";
import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import Add from "../modals/manageaccount/add";
import Edit from "../modals/manageaccount/edit";
import Delete from "../modals/manageaccount/delete";
import { useState, useEffect } from "react";

const ManageCS = () => {
  const [modalEdit, setModalEdit] = useState(false);
  const [modalAdd, setModalAdd] = useState(false);
  const [body, setBody] = useState("");
  const [modalDelete, setModalDelete] = useState(false);
  const [data, setData] = useState(null);

  const getUser = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}getUser`, {
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

  const column = ["#", "Username", "Nama", "Email", "Ubah", "Hapus"];

  useEffect(() => {
    getUser();
  }, []);

  return (
    <div className={style.body}>
      <h3>Manage Akun</h3>
      <br />
      <button
        className={style.buttonHijau}
        onClick={() => {
          setModalAdd(true);
        }}
      >
        + Tambah Data
      </button>
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
              <td>{item.username}</td>
              <td>{item.name}</td>
              <td>{item.email}</td>
              <td>
                <button
                  name="edit"
                  className={style.buttonChange}
                  onClick={() => {
                    setModalEdit(true);
                    setBody(item);
                  }}
                >
                  <CreateIcon sx={{ color: "#B2C154" }} />
                </button>
              </td>
              <td>
                <button
                  name="delete"
                  className={style.buttonChange}
                  onClick={() => {
                    setModalDelete(true);
                    setBody(item);
                  }}
                >
                  <DeleteForeverIcon sx={{ color: "#CC100F" }} />
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <Add show={modalAdd} close={() => {setModalAdd(false); getUser();}} />
      <Edit show={modalEdit} close={() => {setModalEdit(false);getUser();}} data={body} />
      <Delete show={modalDelete} close={() => {setModalDelete(false);getUser();}} data={body} />
    </div>
  );
};

export default ManageCS;
