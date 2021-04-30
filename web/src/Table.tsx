import React from "react"
import styled from "styled-components"

type ITableProps = {
    headerColumns: string[]
    rows: React.ReactNode[][]
}

const TableStyle = styled.table`
    border: 0px;
`

const TableData = styled.td`
    padding: 5px;
    border: 0px;
`

const TableRow = styled.tr<{ striped?: boolean }>`
    margin: 0 5px;
    border: none;
    background : ${props => props.striped ? "#333" : "#444"}
`

export default function Table({ headerColumns, rows }: ITableProps) {
    return <TableStyle>
        <thead>
            <TableRow>
                {headerColumns.map((column, i) => <th key={i}>{column}</th>)}
            </TableRow>
        </thead>
        <tbody>
            {rows.map((row, i) => 
                <TableRow key={i} striped={i % 2 === 0}>
                    {row.map((column, i) => <TableData key={i}>{column}</TableData>)}
                </TableRow>
            )}
        </tbody>
    </TableStyle>
}

