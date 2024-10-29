import { createGlobalStyle } from "styled-components";
import "bootstrap/dist/css/bootstrap.min.css";

export default createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    input:-webkit-autofill,
    input:-webkit-autofill:hover,
    input:-webkit-autofill:focus,
    input:-webkit-autofill:active {
    -webkit-box-shadow: 0 0 0 30px #2a2b2e inset !important;
    }
    input:-webkit-autofill{
    -webkit-text-fill-color: white !important;
    }
  }
  html{
   scroll-behavior: smooth;
   @media (max-width: 1080px) {
      font-size: 93.75%;
    }
    @media (max-width: 720px) {
      font-size: 87.5%;
    }
  }
  body {
    -webkit-font-smoothing: antialiased;
    background: #090a10;
    color: #fff;
  }
  html, body, #__next{
    // height: 100%;
    height: auto!important;
    min-height: 100%;
  }
  body, input, button, textarea {
    font-family: 'Roboto', sans-serif;
  }
  button{
    cursor: pointer;
    outline: none;
    &:focus {
      outline: transparent !important;
    }
  }
  [disabled] {
    opacity: 0.5;
    cursor: not-allowed;
  }
  ul {
    list-style: none;
  }
  a {
    text-decoration: none !important;
  }
  .react-modal-overlay {
    background: rgba(0, 0, 0, 0.6);
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    z-index: 20;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 1rem;
  }
  .react-modal-content {
    max-width: 879px;
    width:100%;
    background: linear-gradient(34.05deg, #101114 1.01%, #454652 100%);
    box-shadow: 0px 4px 13px 1px rgba(0, 0, 0, 0.08);
    padding: 2rem 4rem;
    border-radius: 2rem;
    position: relative;
    @media (max-width: 768px){
      padding:2rem;
    }
  }
`;