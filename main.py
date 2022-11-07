import pandas as pd
import plotly.graph_objects as go

df = pd.read_csv('spy.csv')
fig = go.Figure()
fig.add_trace(go.Candlestick(
                x=df['datetime'],
                open=df['open'],
                high=df['high'],
                low=df['low'],
                close=df['close'],
                increasing_line_color= 'red', decreasing_line_color= 'gray'
                )
             )
fig.update_xaxes(rangeslider_visible=False)
fig.show()
