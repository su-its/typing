'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button, Modal, ModalOverlay, ModalContent, ModalHeader, ModalFooter, ModalBody, ModalCloseButton, useDisclosure, Input } from '@chakra-ui/react';

const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [studentId, setStudentId] = useState(''); // 学籍番号を管理するステート

  const handleRouteGame = async () => {
    // 学籍番号を使用したログイン処理など
    console.log(studentId); // 実際のアプリケーションではここでバックエンドに送信など
    onClose(); // モーダルを閉じる
    router.push("/game");
  };

  return (
    <>
      <Button colorScheme="green" size="lg" onClick={onOpen}>
        Game Start
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>ログイン</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            {/* 学籍番号入力フォーム */}
            <Input 
              placeholder="学籍番号を入力してください" 
              value={studentId}
              onChange={(e) => setStudentId(e.target.value)}
            />
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={handleRouteGame}>
              ログインして続行
            </Button>
            <Button variant="ghost" onClick={onClose}>閉じる</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default GameStartButton;
