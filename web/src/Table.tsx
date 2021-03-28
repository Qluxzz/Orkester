import React from "react"
import styled from "styled-components"

type ITableProps = {
    headerColumns: string[]
    rows: React.ReactNode[][]
}

export function Table({ headerColumns, rows }: ITableProps) {
    return <table>
        <thead>
            <tr>
                {headerColumns.map((column, i) => <th key={i}>{column}</th>)}
            </tr>
        </thead>
        <tbody>
            {rows.map((row, i) => 
                <tr key={i}>
                    {row.map((column, i) => <td key={i}>{column}</td>)}
                </tr>
            )}
        </tbody>
    </table>
}

export const StripedTable = styled(Table)`
    td:nth-child(odd) {
        background: #ddd;
    }
`