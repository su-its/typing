'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button, Modal, ModalOverlay, ModalContent, ModalHeader, ModalFooter, ModalBody, ModalCloseButton, useDisclosure } from '@chakra-ui/react';

const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const handleRouteGame = async () => {
    // TODO:ここでログインの処理を行う
    // ログイン成功後にゲームページへ遷移
    onClose(); 
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
          <ModalHeader>続けるにはログインが必要です</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            {/* ログインフォームをここに配置 */}
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
