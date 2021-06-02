import React from "react"
import { Link } from "react-router-dom"
import styled from "styled-components"

const Container = styled.aside`
    background: #333;
    padding: 10px;
`


export default function SideBar() {
    return <Container>
        <Link to="/collections/tracks">Liked tracks</Link>
    </Container>
}