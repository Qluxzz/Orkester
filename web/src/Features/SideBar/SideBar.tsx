import React from "react"
import { Link } from "react-router-dom"
import styled from "styled-components"

const Container = styled.ul`
    background: #333;
    padding: 10px;
    margin: 0;

    li {
        list-style: none;
        padding: 10px;
    }
`


export default function SideBar() {
    return <Container>
        <Link to="/collection/tracks"><li>Liked tracks</li></Link>
    </Container>
}