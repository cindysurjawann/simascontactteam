import { Button, Modal, ModalBody, ModalFooter } from "reactstrap";
import style from "./modalwadanzoom.module.scss";
const ConfirmationModal = (props) => {
  
  return (
    <>
      <Modal isOpen={props.show} cancel={props.close}>
        <div className="modal-header" style={{ backgroundColor: "#36506A" }}>
          <h5 className="modal-title" id="exampleModalLabel" style={{ color: "white" }}>
            Konfirmasi Perubahan
          </h5>
          <button aria-label="Close" className=" close" type="button" onClick={props.close}>
            <span aria-hidden={true}>×</span>
          </button>
        </div>
        <ModalBody>
          <div className={style.body}>Apakah kamu yakin ingin mengubah link {props.data.newlink} ?</div>
        </ModalBody>
        <ModalFooter>
          <Button className={style.setuju} onClick={props.response} type="button">
            YA
          </Button>
          <Button className={style.tidak} type="button" onClick={props.close}>
            Tidak
          </Button>
        </ModalFooter>
      </Modal>
    </>
  );
};

export default ConfirmationModal;
