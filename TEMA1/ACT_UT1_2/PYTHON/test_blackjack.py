import unittest
from blackjack import getHandValue

class TestBlackjack(unittest.TestCase):

    def test_mano_estandar(self):
        # Test 1: PASA (10 + Jota = 20 puntos)
        mano = [('10', '♥'), ('J', '♠')]
        self.assertEqual(getHandValue(mano), 20)

    def test_mano_con_as(self):
        # Test 2: PASA (Evalúa As como 11 si no pasa de 21, o como 1 si se pasa)
        mano_as_alto = [('A', '♥'), ('9', '♦')]
        self.assertEqual(getHandValue(mano_as_alto), 20)

        mano_as_bajo = [('A', '♥'), ('10', '♠'), ('K', '♦')]
        self.assertEqual(getHandValue(mano_as_bajo), 21)

    def test_calculo_fallido(self):
        # Test 3: FALLA intencionadamente
        # Rey + Reina suman 20, pero afirmamos que es 21
        mano = [('K', '♥'), ('Q', '♠')]
        self.assertEqual(getHandValue(mano), 21)

if __name__ == '__main__':
    unittest.main()
