#!/bin/bash -v
echo ""
echo "num dict example:"
~/go/bin/fcgg -fgg -dictpass -- num_dict.fgg
echo ""
echo "The foo bar subtyping example:"
~/go/bin/fcgg -fgg -dictpass -- stephen_counter_ex.fgg
echo ""
echo "test the difference of parameter names in different decls:"
~/go/bin/fcgg -fgg -monom -- type_param_order.fgg
~/go/bin/fcgg -fgg -dictpass -- type_param_order.fgg
