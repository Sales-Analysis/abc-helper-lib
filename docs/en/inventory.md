# Inventory Analyses

This page is the index for inventory-focused analyses implemented in
`abc-helper-lib`.

## Table Of Contents

- [ABC](inventory/abc.md)
- [XYZ](inventory/xyz.md)
- [ABC-XYZ](inventory/abcxyz.md)
- [VED](inventory/ved.md)
- [FSN](inventory/fsn.md)
- [HML](inventory/hml.md)
- [SDE](inventory/sde.md)
- [EOQ](inventory/eoq.md)
- [Reorder Point](inventory/reorderpoint.md)
- [Safety Stock](inventory/safetystock.md)
- [Pareto 80/20](inventory/pareto.md)
- [GM/Contribution](inventory/gmcontribution.md)
- [Service Level](inventory/servicelevel.md)
- [Decision guide](decision-guide.md)
- [API conventions](api.md)

## Summary Table

| Analysis | Package | Facade Method | Primary Inputs | Main Output |
| :--- | :--- | :--- | :--- | :--- |
| [`ABC`](inventory/abc.md) | `abc` | `Inventory().ABC` | quantity, unit price | revenue share and group |
| [`XYZ`](inventory/xyz.md) | `xyz` | `Inventory().XYZ` | demand series | demand variability group |
| [`ABC-XYZ`](inventory/abcxyz.md) | `abcxyz` | `Inventory().ABCXYZ` | sales value and demand series | combined group |
| [`VED`](inventory/ved.md) | `ved` | `Inventory().VED` | criticality score | criticality group |
| [`FSN`](inventory/fsn.md) | `fsn` | `Inventory().FSN` | movement history | fast/slow/non-moving group |
| [`HML`](inventory/hml.md) | `hml` | `Inventory().HML` | unit cost | cost-based group |
| [`SDE`](inventory/sde.md) | `sde` | `Inventory().SDE` | lead time in days | procurement difficulty group |
| [`EOQ`](inventory/eoq.md) | `eoq` | `Inventory().EOQ` | demand, ordering cost, holding cost | optimal order quantity |
| [`Reorder Point`](inventory/reorderpoint.md) | `reorderpoint` | `Inventory().ReorderPoint` | average demand, lead time, safety stock | reorder point |
| [`Safety Stock`](inventory/safetystock.md) | `safetystock` | `Inventory().SafetyStock` | demand deviation, lead time, service factor | safety stock |
| [`Pareto 80/20`](inventory/pareto.md) | `pareto` | `Inventory().Pareto` | value per item | cumulative value and 80/20 check |
| [`GM/Contribution`](inventory/gmcontribution.md) | `gmcontribution` | `Inventory().GMContribution` | revenue and cost components | margin metrics |
| [`Service Level`](inventory/servicelevel.md) | `servicelevel` | `Inventory().ServiceLevel` | demand fulfillment and stockout cycles | service level metrics |

## Related Technical Notes

- [Decision guide](decision-guide.md)
- [API conventions](api.md)
- `ABC`, `XYZ`, and `ABC-XYZ` share input alignment through `OriginalIndex`
- `Safety Stock` and `Reorder Point` are operationally related and are usually
  read together
