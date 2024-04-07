"use client";

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
import { showWarningToast } from "@/utils/toast";

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
  state: any;
  dispatchAction: (payload: FormData) => void;
  pending: boolean;
}

const LoginModalPresenter: React.FC<LoginModalProps> = ({ isOpen, onClose, state, dispatchAction, pending }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <form action={ async (formData: FormData) => {
          await dispatchAction(formData);
          state.error && showWarningToast(state.error);
        }}>
          <ModalHeader>続けるにはログインが必要です</ModalHeader>
          <ModalBody>
            <Input
              isRequired
              required
              name="student-number"
              placeholder="学籍番号を入力してください"
              pattern="[0-9A-Z]{8}"
              title="学籍番号"
              role="textbox"
            />
          </ModalBody>

          <ModalFooter>
            <Button type="submit" colorScheme="blue" mr={3} isLoading={pending} role="submit">
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

export { LoginModalPresenter };

interface LoginModalContainerProps {
  isOpen: boolean;
  onClose: () => void;
}

const LoginModalContainer: React.FC<LoginModalContainerProps> = ({ isOpen, onClose }) => {
  const [state, dispatchAction, pending] = useFormState(login, {});

  return (
    <LoginModalPresenter
      isOpen={isOpen}
      onClose={onClose}
      state={state}
      dispatchAction={dispatchAction}
      pending={pending}
    />
  );
};

export default LoginModalContainer;
