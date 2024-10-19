import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect"; // para matchers como `toBeInTheDocument`
import Home from "."; // importa el componente Home
import { validateEmail } from "@/lib/utils/email-validation";
import { describe, expect, it, jest } from "@jest/globals";

// Mock de la función `validateEmail`
jest.mock("@/lib/utils/email-validation", () => ({
  validateEmail: jest.fn(),
}));

describe("Home Component", () => {
  const mockProps = {
    topics: ["React", "Node.js", "Vue"],
    email: "",
    selectedTopics: [],
    onChange: jest.fn(),
    handleSubscribe: jest.fn(),
    onSetEmail: jest.fn(),
    t: jest.fn<(key: string) => string>().mockImplementation((key) => key),
  };

  describe("should render the component correctly", () => {
    render(<Home {...mockProps} />);

    // Verifica si el título se renderiza correctamente
    expect(
      screen.getByText("Suscríbete a nuestro boletin")
    ).toBeInTheDocument();
  });

  it("should call onSetEmail when the email input changes", () => {
    render(<Home {...mockProps} />);

    // Verifica si los checkboxes de los topics se renderizan correctamente
    expect(screen.getByLabelText("React")).toBeInTheDocument();
    expect(screen.getByLabelText("Node.js")).toBeInTheDocument();
    expect(screen.getByLabelText("Vue")).toBeInTheDocument();

    // Simula el cambio de valor en el campo de email
    const emailInput = screen.getByPlaceholderText("example@example.com");
    fireEvent.change(emailInput, { target: { value: "test@example.com" } });

    // Verifica que `onSetEmail` haya sido llamado con el valor correcto
    expect(mockProps.onSetEmail).toHaveBeenCalledWith("test@example.com");
  });

  it("should disable the submit button if no topics are selected or the email is invalid", () => {
    // Simular la validación de email
    (validateEmail as jest.Mock).mockReturnValue(false);

    render(<Home {...mockProps} />);

    // Verifica que el botón de submit esté deshabilitado cuando no hay topics seleccionados
    const submitButton = screen.getByText("subscribe");
    expect(submitButton).toBeDisabled();
  });

  it("should enable the submit button if there are selected topics and the email is valid", () => {
    // Simular un email válido y temas seleccionados
    (validateEmail as jest.Mock).mockReturnValue(true);

    const propsWithSelectedTopics = {
      ...mockProps,
      email: "test@example.com",
      selectedTopics: ["React"],
    };

    render(<Home {...propsWithSelectedTopics} />);

    // Verifica que el botón de submit esté habilitado cuando hay un email válido y temas seleccionados
    const submitButton = screen.getByText("subscribe");
    expect(submitButton).not.toBeDisabled();
  });

  it("should call handleSubscribe when the submit button is clicked", () => {
    (validateEmail as jest.Mock).mockReturnValue(true);

    const propsWithSelectedTopics = {
      ...mockProps,
      email: "test@example.com",
      selectedTopics: ["React"],
    };

    render(<Home {...propsWithSelectedTopics} />);

    const submitButton = screen.getByText("subscribe");
    fireEvent.click(submitButton);

    // Verifica que la función `handleSubscribe` haya sido llamada cuando se hace clic en el botón
    expect(mockProps.handleSubscribe).toHaveBeenCalled();
  });

  it("should call onChange when a topic is selected", () => {
    render(<Home {...mockProps} />);

    const checkbox = screen.getByLabelText("React");
    fireEvent.click(checkbox);

    // Verifica que `onChange` haya sido llamado al seleccionar un checkbox
    expect(mockProps.onChange).toHaveBeenCalled();
  });
});
