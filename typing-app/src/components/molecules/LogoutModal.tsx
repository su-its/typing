import React from "react";
import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, ModalFooter, Button } from "@chakra-ui/react";

interface LogoutModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
}

const LogoutModal: React.FC<LogoutModalProps> = ({ isOpen, onClose, onConfirm }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>ログアウトしました</ModalHeader>
        <ModalBody>ご利用ありがとうございました。</ModalBody>
        <ModalFooter>
          <Button colorScheme="blue" mr={3} onClick={onConfirm}>
            OK
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default LogoutModal;
