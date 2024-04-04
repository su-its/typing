import React from "react";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  Button,
  Input,
} from "@chakra-ui/react";
import { login } from "@/app/actions";
import { useFormState } from "react-dom";

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
}

const LoginModal: React.FC<LoginModalProps> = ({ isOpen, onClose }) => {
  const [state, dispatchAction, pending] = useFormState(login, {});
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <form action={dispatchAction}>
          <ModalHeader>続けるにはログインが必要です</ModalHeader>
          <ModalBody>
            <Input
              isRequired
              required
              name="student-number"
              placeholder="学籍番号を入力してください"
              pattern="[0-9A-Z]{8}"
              title="学籍番号"
            />
            {state.error ? `エラー: ${state.error}` : null}
          </ModalBody>

          <ModalFooter>
            <Button type="submit" colorScheme="blue" mr={3} isLoading={pending}>
              ログインして続行
            </Button>
            <Button variant="ghost" onClick={onClose}>
              閉じる
            </Button>
          </ModalFooter>
        </form>
      </ModalContent>
    </Modal>
  );
};

export default LoginModal;
