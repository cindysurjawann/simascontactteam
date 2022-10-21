import React,{ useState, useEffect } from "react";
import style from "./managewa.module.scss";
import ConfirmationModal from "../modals/modalwadanzoom";


const ManageWa = () => {
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  useEffect(() => {
    getWa();
  }, []);

  const [modalOpen, setModalOpen] = useState(false);
  const [body, setBodyData] = useState("");

  const onSubmit = async (e) => {
    const dataform = {
      newlink: newLink,
    };
    setBodyData(dataform);
    setModalOpen(true);
  };
  async function getWa(e) {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL}getlink?linktype=WA`, {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    });
    const data = await res.json();
    setData(data);
    console.log(data);
  }
  async function putWa() {
    setModalOpen(false);

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}updatelink?linktype=WA`, {
        method: "PUT",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({
          linkvalue: body.newlink,
          UpdatedBy: "system",
        }),
      });
      if(res.status != 200){
        throw "gagal mendapatkan pesan WA"();
      }
      const d = { ...data };
      d.data.linkvalue = body.newlink;
      console.log(d);
      setData(d);
      alert("Update Sukses");
    } catch (error) {
      alert("Update Gagal");
    }
  }

  return (
    <div className={style.zoom}>
      <h1>Manage Link WhatsApp</h1>
      <div className={style.inputbox}>
        <form
          onSubmit={(e) => {
            e.preventDefault();
          }}
        >
          <div>
            <h3>Link WhatsApp Lama</h3>
            <input className={style.readonly} type="text" placeholder={!data ? "" : data.data.linkvalue} readOnly disabled="true" />
          </div>
          <br />
          <div>
            <h3>Link WhatsApp Baru</h3>
            <input type="text" name="newlink" required value={newLink} onChange={(e) => setNewLink(e.target.value)} />
          </div>
          <br />
          <br />
          <button className={style.buttonHijau} onClick={onSubmit}>
            SIMPAN
          </button>
        </form>
      </div>
      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} linktype={"Wa"} data={body} response={putWa} />;
    </div>
  );
};

export default ManageWa;
