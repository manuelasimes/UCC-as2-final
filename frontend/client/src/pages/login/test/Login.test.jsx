import { describe, it, expect } from "vitest";
import { render, screen, fireEvent } from '@testing-library/react'
import matchers from "@testing-library/jest-dom/matchers";
import Login from '../Login';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";

expect.extend(matchers);

describe("Test del componente Login", () => {
    it("Se renderiza el componente en el dom", () => {
      render(
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Login/>}/>
        </Routes> 
      </BrowserRouter>
      )
      const EmailInput = screen.getByPlaceholderText("Ingrese su correo");
      expect(EmailInput).toBeDefined();
      expect(EmailInput.value).toEqual("");
      fireEvent.change(EmailInput, { target: { value: "marianapaulina@dh.com" } })
      expect(EmailInput.value).toEqual("marianapaulina@dh.com")
    })
  })