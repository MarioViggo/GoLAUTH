import styled from "styled-components";

export const Container = styled.div`
    display: flex;
    position: relative;
    padding: 1rem 0;
    flex-direction: column;
    text-align: start;
    label {
        /* position: absolute; */
        /* top: .5rem; */
        /* left: 1rem; */
        font-size: 0.9rem;
        color: #fff;
        font-weight: 700;
        /* background: #37363c;    */
    }
    input {
        background: #37363c;
        border: 1px solid #fff;
        color: #fff;
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;
        outline: none;
        transition: border 0.5s;

        &:focus, &:hover {
            border: 1px solid rgba(255, 255, 255, 0.85);
        }
    }
`;