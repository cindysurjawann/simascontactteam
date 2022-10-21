import { Button, Modal, ModalBody } from "reactstrap";
import style from "./modalDetailCart.module.scss";
const ConfirmationDetailCart = (props) => {
  const snk = props.data.syarat.split(";");
  const format = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
  });
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
            <p>
                <h4> Premi </h4>
                {format.format(props.data.premi)}
            </p>

            <p>
                <h4> Uang Pertanggungan </h4>
                {format.format(props.data.uangpertanggungan)}
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
