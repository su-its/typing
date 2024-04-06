import { createStandaloneToast } from '@chakra-ui/react';

const { toast: toastChakra } = createStandaloneToast();

const toast = ({
  status,
  title = '',
  description = '',
}: {
  status: 'success' | 'error' | 'warning' | 'info';
  title?: string;
  description?: string;
}) => {
  toastChakra({
    title,
    description,
    status,
    duration: 3000,
    isClosable: true,
    position: 'top',
  });
};

export const showSuccessToast = (title: string, description?: string) => {
  toast({
    title,
    description,
    status: 'success',
  });
};

export const showWarningToast = (title: string, description?: string) => {
  toast({
    title,
    description,
    status: 'warning',
  });
};

export const showErrorToast = (title: string, description?: string) => {
  toast({
    title,
    description,
    status: 'error',
  });
};

export default toast;