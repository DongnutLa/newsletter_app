import { describe, expect, test } from "@jest/globals";
import { validateEmail } from "./email-validation";

describe("validateEmail", () => {
  test("Correo inválido 1", () => {
    const result = validateEmail("aaaaa");
    expect(result).toBe(false);
  });

  test("Correo inválido 2", () => {
    const result = validateEmail("example@example");
    expect(result).toBe(false);
  });

  test("debe retornar true", () => {
    const result = validateEmail("example@example.com");
    expect(result).toBe(true);
  });
});
