This plugin intended to Add Client request path as Http Header(`X-Forwarded-Path`) why?

For example you want to see what's slowest performing path in your application from prometheus metrics. Then this plugin comes handy.

```yaml
prometheus:
  headerLabels:
    path: X-Forwarded-Path
```