import styled from 'styled-components';

export const Container = styled.main`
    padding: 2rem 0;
    min-height: 100vh;
`;

export const Card = styled.div`
    border: 1px solid red;
    padding: 1.5rem 1.5rem 3rem;
    margin-bottom: 1rem;
    border-radius: 8px;
    min-height: 12rem;
    position: relative;

    &:hover {
        cursor: pointer;
        box-shadow: 0 0 10px #fff;
    }
    
    h4 {
        margin-bottom: 1rem;
    }

    .arrow {
        position: absolute;
        bottom: 1rem;
        right: 1rem;
        display: flex;
        justify-content: flex-end;
    }
`;