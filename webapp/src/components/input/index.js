import React from "react";
import { Container } from "./styles";

function Input ({name, label, type, value, onChange, placeholder, required, disabled, className, style, id}) {
    return (
        <Container>
            <label htmlFor={name}>{label}</label>
            <input
                id={id}
                name={name}
                type={type}
                value={value}
                onChange={onChange}
                placeholder={placeholder}
                required={required}
                disabled={disabled}
                className={className}
                style={style}
            />
        </Container>
    )
};

export default Input;