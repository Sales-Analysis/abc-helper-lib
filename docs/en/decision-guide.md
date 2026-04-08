# Analysis Selection Guide

Use this page when the main question is not "how does the package work?" but
"which analysis should I choose first?"

## Inventory Questions

| Business question | Start with | Often pair with |
| :--- | :--- | :--- |
| Which SKUs drive most revenue? | [ABC](inventory/abc.md) | [Pareto 80/20](inventory/pareto.md), [ABC-XYZ](inventory/abcxyz.md) |
| Do a small number of items generate most of the value? | [Pareto 80/20](inventory/pareto.md) | [ABC](inventory/abc.md) |
| Which items have stable vs volatile demand? | [XYZ](inventory/xyz.md) | [FSN](inventory/fsn.md), [ABC-XYZ](inventory/abcxyz.md) |
| I need both value importance and demand stability. | [ABC-XYZ](inventory/abcxyz.md) | [Safety Stock](inventory/safetystock.md), [Reorder Point](inventory/reorderpoint.md) |
| Which items are operationally critical? | [VED](inventory/ved.md) | [SDE](inventory/sde.md), [HML](inventory/hml.md) |
| Which items move fast, slow, or not at all? | [FSN](inventory/fsn.md) | [XYZ](inventory/xyz.md), [Pareto 80/20](inventory/pareto.md) |
| Which items are expensive per unit? | [HML](inventory/hml.md) | [ABC](inventory/abc.md), [VED](inventory/ved.md) |
| Which items are hard to source? | [SDE](inventory/sde.md) | [Safety Stock](inventory/safetystock.md), [VED](inventory/ved.md) |
| What order quantity is economically optimal? | [EOQ](inventory/eoq.md) | [Reorder Point](inventory/reorderpoint.md) |
| When should I place the next order? | [Reorder Point](inventory/reorderpoint.md) | [Safety Stock](inventory/safetystock.md), [Service Level](inventory/servicelevel.md) |
| How much protective buffer stock do I need? | [Safety Stock](inventory/safetystock.md) | [Reorder Point](inventory/reorderpoint.md), [Service Level](inventory/servicelevel.md) |
| How well do we fulfill demand today? | [Service Level](inventory/servicelevel.md) | [Safety Stock](inventory/safetystock.md), [Reorder Point](inventory/reorderpoint.md) |
| Which items create gross margin or contribution? | [GM/Contribution](inventory/gmcontribution.md) | [ABC](inventory/abc.md), [Pareto 80/20](inventory/pareto.md) |

## Customer Questions

| Business question | Start with | Often pair with |
| :--- | :--- | :--- |
| Which customers are most engaged right now? | [RFM](customer/rfm.md) | [CLV](customer/clv.md), [Churn / Retention](customer/churn.md) |
| Which customers are economically most valuable? | [CLV](customer/clv.md) | [RFM](customer/rfm.md), [Churn / Retention](customer/churn.md) |
| Which customers are likely retained, at risk, or churned? | [Churn / Retention](customer/churn.md) | [RFM](customer/rfm.md), [CLV](customer/clv.md) |

## Recommended Combinations

- `ABC + XYZ + Safety Stock`: prioritize high-value unstable items and size their protection.
- `VED + SDE + Safety Stock`: protect critical items that are also difficult to replenish.
- `ABC + Pareto`: validate whether a small SKU share drives most revenue.
- `RFM + CLV`: combine engagement and economic value before campaign decisions.
- `Churn + CLV`: focus retention effort where potential loss is largest.

## Quick Heuristics

- Use [ABC](inventory/abc.md) when the language is "value", "revenue", or "contribution share".
- Use [XYZ](inventory/xyz.md) when the language is "stability", "variability", or "forecastability".
- Use [FSN](inventory/fsn.md) when the language is "movement frequency", not statistical variability.
- Use [VED](inventory/ved.md) when business criticality matters more than pure economics.
- Use [EOQ](inventory/eoq.md), [Safety Stock](inventory/safetystock.md), and [Reorder Point](inventory/reorderpoint.md) as an operating trio instead of isolated formulas.

## Common Selection Mistakes

- Using [ABC](inventory/abc.md) to infer demand stability. That is usually a [XYZ](inventory/xyz.md) question.
- Using [XYZ](inventory/xyz.md) as a proxy for business importance. That is usually an [ABC](inventory/abc.md) or [Pareto 80/20](inventory/pareto.md) question.
- Using [EOQ](inventory/eoq.md) without checking reorder timing and service assumptions.
- Using [CLV](customer/clv.md) alone for campaign prioritization without a current-state model such as [RFM](customer/rfm.md) or [Churn](customer/churn.md).
