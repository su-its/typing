"use client";

import React from "react";
import { Button, useDisclosure } from "@chakra-ui/react";
import LogoutModal from "./LogoutModal"; 
import { logout } from "@/app/actions";

const LogoutButton: React.FC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  const handleLogout = async () => {
    await logout();
    onOpen();
  };

  return (
    <>
      <Button colorScheme="blue" size="lg" onClick={handleLogout}>
        Logout
      </Button>

      <LogoutModal isOpen={isOpen} onClose={onClose} />
    </>
  );
};

export default LogoutButton;
