import style from "./managezoom.module.scss";
import ConfirmationModal from "../modals/modalwadanzoom";
import React, { useState, useEffect } from "react";

const ManageWa = () => {
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  useEffect(() => {
    getZoom();
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
  async function getZoom(e) {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL}getlink?linktype=Zoom`, {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    });
    const data = await res.json();
    setData(data);
    console.log(data);
  }
  async function putZoom() {
    setModalOpen(false);

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}updatelink?linktype=Zoom`, {
        method: "PUT",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({
          linkvalue: body.newlink,
          UpdatedBy: "system",
        }),
      });
      if (res.status != 200) {
        throw "gagal mendapatkan pesan Zoom"();
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
      <h1>Manage Link Zoom</h1>
      <div className={style.inputbox}>
        <form
          onSubmit={(e) => {
            e.preventDefault();
          }}
        >
          <div>
            <h3>Link Zoom Lama</h3>
            <input className={style.readonly} type="text" placeholder={!data ? "" : data.data.linkvalue} readOnly disabled="true" />
          </div>
          <br />
          <div>
            <h3>Link Zoom Baru</h3>
            <input type="text" name="newlink" required value={newLink} onChange={(e) => setNewLink(e.target.value)} />
          </div>
          <br />
          <br />
          <button className={style.buttonHijau} onClick={onSubmit}>
            SIMPAN
          </button>
        </form>
      </div>
      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} linktype={"Zoom"} data={body} response={putZoom} />;
    </div>
  );
};

export default ManageWa;
