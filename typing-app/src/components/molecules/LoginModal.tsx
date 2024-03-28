import React, { useState } from "react";
import { useRouter } from "next/navigation";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Button,
  Input,
} from "@chakra-ui/react";

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
}

const LoginModal: React.FC<LoginModalProps> = ({ isOpen, onClose }) => {
  const [studentId, setStudentId] = useState("");
  const router = useRouter();

  const handleLogin = async () => {
    // TODO:ログイン済みかどうかを判別
    // TODO:学籍番号を使用したログイン処理
    console.log(studentId);
    router.push("/game");
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>続けるにはログインが必要です</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <Input
            placeholder="学籍番号を入力してください"
            value={studentId}
            onChange={(e) => setStudentId(e.target.value)}
          />
        </ModalBody>

        <ModalFooter>
          <Button colorScheme="blue" mr={3} onClick={handleLogin}>
            ログインして続行
          </Button>
          <Button variant="ghost" onClick={onClose}>
            閉じる
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default LoginModal;
