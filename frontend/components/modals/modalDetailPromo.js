import { Button, Modal, ModalBody } from "reactstrap";
import style from "./modalDetailCart.module.scss";
const ConfirmationDetailCart = (props) => {
  const snk = props.data.syarat.split(";");
  return (
    <>
      <div className={style.container}>
        <Modal
          className={style.container}
          isOpen={props.show}
          cancel={props.close}
          toggle={props.close}
        >
          <div className="modal-header" style={{ backgroundColor: "white" }}>
            <h3 className="modal-title" id="exampleModalLabel">
              {props.data.judul}
            </h3>
            <br />
            <Button
              aria-label="Close"
              className=" close"
              type="button"
              onClick={props.close}
            >
              <span aria-hidden={true}>×</span>
            </Button>
          </div>
          <ModalBody>
            <div className={style.body}>
              <img
                className={style.cslaki}
                src={props.data.foto}
                alt="cslaki"
              />
            </div>
            <p className={style.periode}>
              Periode: {props.data.startdate.substring(0, 10)} s/d{" "}
              {props.data.enddate.substring(0, 10)}
            </p>
            <h4>Deskripsi</h4>
            <div className={style.deskripsi}>{props.data.deskripsi}</div>
            <br />
            <h4>Syarat dan Ketentuan</h4>
            {snk.map((item, index) => (
              <p key={index}>{item}</p>
            ))}
          </ModalBody>
        </Modal>
      </div>
    </>
  );
};

export default ConfirmationDetailCart;
