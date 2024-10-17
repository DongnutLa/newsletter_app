import styled from "styled-components";

export const LoginTitle = styled.h1`
  color: #f53954;
  text-align: center;
`;

export const Form = styled.form`
  display: flex;
  gap: 20px;
  flex-direction: column;
  justify-content: center;
  background-color: white;
  padding: 20px;
  border-radius: 20px;
`;

export const LoginWrapper = styled.div`
  display: grid;
  place-content: center;
  height: 100vh;
  width: 100vw;
  background-color: #f53954;
  gap: 20px;
`;

export const LoginInputWrapper = styled.div`
  background-color: #f53954;
  display: flex;
  justify-content: space-between;
  padding: 4px;
  height: 40px;
  border-radius: 40px;
`;

export const LoginInput = styled.input`
  border: unset;
  border-radius: 40px;
  &:focus-visible {
    outline: unset;
  }
  margin-left: 12px;
  width: 100%;
  background-color: #f53954;
  color: white;

  &::placeholder {
    color: white;
  }
`;

export const LoginSubmit = styled.button`
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

export const ErrorFeedback = styled.span`
  color: #f53954;
  font-size: 14px;
  padding-left: 20px;
`;
