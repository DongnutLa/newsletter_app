import styled, { css } from "styled-components";

export const NewsletterFormWrapper = styled.main`
  display: grid;
  place-content: center;
  height: 100vh;
  width: 100vw;
  gap: 20px;
`;

export const Form = styled.form`
  display: flex;
  gap: 12px;
  flex-direction: column;
  justify-content: center;
  background-color: white;
  padding: 20px;
  border-radius: 20px;
`;

export const FormTitle = styled.h1`
  text-align: center;
`;

export const InputWrapper = styled.div<{ textarea?: boolean }>`
  background-color: #c7c7c7;
  display: flex;
  justify-content: space-between;
  padding: 4px;
  height: 40px;
  width: 350px;
  border-radius: 12px;

  ${({ textarea }) =>
    textarea &&
    css`
      height: 150px;
    `}
`;

export const Input = styled.input`
  border: unset;
  border-radius: 12px;
  &:focus-visible {
    outline: unset;
  }

  width: 100%;
  background-color: #c7c7c7;

  &::placeholder {
    color: white;
  }
`;

export const TextArea = styled.textarea`
  border: unset;
  border-radius: 12px;
  &:focus-visible {
    outline: unset;
  }

  width: 100%;
  height: 150px;
  background-color: #c7c7c7;

  &::placeholder {
    color: white;
  }
`;

export const ErrorFeedback = styled.span`
  color: #f53954;
  font-size: 14px;
  padding-left: 20px;
`;

export const NewsletterUploaderWrapper = styled.div`
  display: flex;
  justify-content: center;
`;

export const NewsletterSubmit = styled.button`
  border-radius: 40px;
  color: white;
  font-weight: bold;
  background-color: #f53954;
  border: unset;
  width: 350px;
  height: 40px;
  cursor: pointer;

  &:hover {
    background-color: #c7c7c7;
    transition: 0.3s;
  }
`;
