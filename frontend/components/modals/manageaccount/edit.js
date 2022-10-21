import { Modal, ModalBody } from "reactstrap";
import style from "./edit.module.scss";

const ModalEdit = (props) => {
  const updateUser = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      name: formData.get("nama"),
      role: parseInt(formData.get("role")),
      email: formData.get("email"),
    };
    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_URL}updateuser?username=${props.data.username}`,
        {
          method: "PUT",
          headers: {
            Authorization: localStorage.getItem("token"),
          },
          body: JSON.stringify(body),
        }
      );
      if (res.status != 200) {
        throw new Error("gagal mengubah data user CS");
      }
      alert("Sukses mengubah data");
      props.close();
    } catch (e) {
      alert("Gagal mengubah data user CS, silahkan refresh ulang");
      return false;
    }
  };

  return (
    <>
      <Modal isOpen={props.show} toggle={props.close}>
        <ModalBody>
          <div style={{ padding: "20px" }}>
            <h4
              style={{
                textAlign: "center",
                paddingTop: "10px",
                paddingBottom: "10px",
              }}
            >
              Form Ubah Data
            </h4>
            <br />
            <form onSubmit={updateUser}>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label for="exampleInputEmail1">Username</label>
                <input
                  type="text"
                  className="form-control"
                  name="username"
                  aria-describedby="emailHelp"
                  placeholder="Masukkan username"
                  defaultValue={props.data.username}
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                  readOnly
                />
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Nama</label>
                <input
                  type="text"
                  className="form-control"
                  name="nama"
                  placeholder="Masukkan nama"
                  defaultValue={props.data.name}
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Posisi</label>
                <select
                  name="role"
                  class="form-control"
                  value={props.data.role}
                >
                  <option value="2">Customer Service</option>
                </select>
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Email</label>
                <input
                  type="text"
                  className="form-control"
                  name="email"
                  placeholder="Masukkan nama"
                  defaultValue={props.data.email}
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
              </div>
              <button
                type="submit"
                className={style.buttonHijau}
                style={{ marginTop: "20px" }}
              >
                Kirim
              </button>
            </form>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

export default ModalEdit;
